package http_utils

import "net/http"

func CheckHandlerAndWrap(method string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			_, err := w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
			if err != nil {
				return
			}
		}

		fn(w, r)
	}
}

func WithGET(fn http.HandlerFunc) http.HandlerFunc {
	return CheckHandlerAndWrap(http.MethodGet, fn)
}

func WithPOST(fn http.HandlerFunc) http.HandlerFunc {
	return CheckHandlerAndWrap(http.MethodPost, fn)
}

func WithDelete(fn http.HandlerFunc) http.HandlerFunc {
	return CheckHandlerAndWrap(http.MethodDelete, fn)
}

func WithPut(fn http.HandlerFunc) http.HandlerFunc {
	return CheckHandlerAndWrap(http.MethodPut, fn)
}

func WithDefaultHandlers(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fn(w, r)
	}
}
