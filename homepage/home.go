package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello GopherCon UK 2018!"

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next.ServeHTTP(w, r)
	})
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	finalHandler := http.HandlerFunc(h.Home)
	mux.Handle("/", h.Logger(finalHandler))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
