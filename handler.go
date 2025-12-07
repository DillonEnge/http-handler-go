package handle

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type HandlerFuncWithError func(w http.ResponseWriter, r *http.Request) *HandlerError

func NewHandlerFunc(h HandlerFuncWithError, hfs ...http.HandlerFunc) http.HandlerFunc {
	if len(hfs) == 0 {
		return defaultHandlerFunc(h)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, hf := range hfs {
			hf(w, r)
		}

	}
}

func defaultHandlerFunc(h HandlerFuncWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("encountered error while serving request", "err", err)
			w.WriteHeader(err.Status)
			errJSON := fmt.Sprintf(
				`{"error": "%s"}`,
				strings.ReplaceAll(err.Error(), `"`, `'`),
			)
			w.Write([]byte(errJSON))
		}
	}
}
