package httpx

import (
	"errors"
	"net/http"
)

const xForwardedFor = "X-Forwarded-For"

// GetFormValues returns the form values.
func GetFormValues(r *http.Request) (map[string]any, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	if err := r.ParseMultipartForm(maxMemory); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			return nil, err
		}
	}

	params := make(map[string]any, len(r.Form))
	for name, values := range r.Form {
		filtered := make([]string, 0, len(values))
		for _, v := range values {
			if len(v) > 0 {
				filtered = append(filtered, v)
			}
		}

		if len(filtered) > 0 {
			params[name] = filtered
		}
	}

	return params, nil
}

// GetRemoteAddr returns the peer address, supports X-Forward-For.
func GetRemoteAddr(r *http.Request) string {
	v := r.Header.Get(xForwardedFor)
	if len(v) > 0 {
		return v
	}

	return r.RemoteAddr
}
