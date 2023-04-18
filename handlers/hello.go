package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world!")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Opps!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "The data that you sent was: %s", d)
}
