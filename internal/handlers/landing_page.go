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
		title:       "Lukas Grimm",
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
			profilePicture: imageModel{
				source:      "/static/profile.jpg",
				alternative: "Headshot of Lukas Grimm",
				sources: []imageSource{
					{
						sourceSet: "/static/profile@1x.avif 1x, /static/profile@2x.avif 2x, /static/profile@3x.avif 3x",
						mediaType: "image/avif",
					},
					{
						sourceSet: "/static/profile@1x.webp 1x, /static/profile@2x.webp 2x, /static/profile@3x.webp 3x",
						mediaType: "image/webp",
					},
					{
						sourceSet: "/static/profile@1x.jpg 1x, /static/profile@2x.jpg 2x, /static/profile@3x.jpg 3x",
						mediaType: "image/jpeg",
					},
				},
			},
		})).Render(request.Context(), buffer)
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
