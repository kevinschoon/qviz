package http

import "net/http"

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
	return http.ListenAndServe(opts.Address, mux)
}

func mainHandler(opts *Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
