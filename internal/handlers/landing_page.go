package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type LandingPageController struct {
	Manifest  Manifest
	BuildTime time.Time
}

func (c LandingPageController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	layoutData := layoutModel{
		title: "Lukas Grimm",

		stylesheets: c.Manifest.MapEntrypoints("index.css"),
		scripts:     c.Manifest.MapEntrypoints("index.ts"),
	}

	buffer := new(bytes.Buffer)
	err := layout(
		layoutData, 
		landingPage(landingPageModel{
			links: []landingPageLink{
				{
					label: "GitHub",
					href:  "https://github.com/untanky",
				},
				{
					label: "LinkedIn",
					href:  "https://linkedin.com/in/lukasgrimm",
				},
			},
		})).
		Render(request.Context(), buffer)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	header := writer.Header()
	header.Add("Content-Length", fmt.Sprintf("%d", buffer.Len()))
	header.Add("Content-Type", "text/html; charset=utf-8")
	header.Add("Content-Language", "en")
	header.Add("Last-Modified", c.BuildTime.Format(http.TimeFormat))
	header.Add("Content-Security-Policy", "default-src 'self'; script-src 'self';")
	header.Add("X-Content-Type-Options", "nosniff")
	header.Add("X-Frame-Options", "DENY")
	header.Add("Referrer-Policy", "origin-when-cross-origin")
	writer.WriteHeader(http.StatusOK)
	buffer.WriteTo(writer)
}

