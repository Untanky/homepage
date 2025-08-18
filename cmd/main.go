package main

import (
	"net/http"
	"strings"
)

type staticHandler struct {
	prefix string
	handler http.Handler
}

type MainHandler struct {
	staticHandlers []staticHandler
}

func (handler *MainHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, h := range handler.staticHandlers {
		if s, ok := strings.CutPrefix(request.URL.Path, h.prefix); ok {
			request.URL.Path = s

			h.handler.ServeHTTP(writer, request)
			return
		}
	}

	writer.WriteHeader(200)
	model := layoutModel {
		title: "Lukas Grimm",
	}

	data := landingPageModel {
		links: []landingPageLink {
			{ label: "GitHub", href: "https://github.com/untanky" },
			{ label: "LinkedIn", href: "https://linkedin.com/in/lukasgrimm" },
		},
	}

	layout(model, landingPage(data)).Render(request.Context(), writer)
}

func (handler *MainHandler) RegisterStatic(httpPrefix string, osPath string) {
	fileServer := http.FileServer(http.Dir(osPath))

	handler.staticHandlers = append(handler.staticHandlers, staticHandler{
		prefix: httpPrefix,
		handler: fileServer,
	})
}

func main() {
	myHandler := &MainHandler{}
	
	myHandler.RegisterStatic("/assets", "./assets")
	myHandler.RegisterStatic("/static", "../static")

	server := &http.Server {
		Addr: ":8080",
		Handler: myHandler,
	}
		
	println("Listening on port 8080...")
	err := server.ListenAndServe()
	panic(err)
}

