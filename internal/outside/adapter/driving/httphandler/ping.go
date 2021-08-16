package httphandler

import (
	"net/http"
)

func (h HTTPHandler) ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello"))
}
