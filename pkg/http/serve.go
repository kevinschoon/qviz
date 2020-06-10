package http

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type TemplateContext struct{}

type Options struct {
	Address string `json:"address"`
}

func DefaultOptions() *Options {
	return &Options{
		Address: "127.0.0.1:3232",
	}
}

func Serve(opts *Options) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler(opts))
	mux.HandleFunc("/svg", svgHandler(opts))
	return http.ListenAndServe(opts.Address, mux)
}

func mainHandler(opts *Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, _ := ioutil.ReadFile("./static/index.html")
		tmpl, _ := template.New("index").Parse(string(raw))
		tmpl.Execute(w, &TemplateContext{})
	}
}

func svgHandler(opts *Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
	}
}
