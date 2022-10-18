package handler

import (
	"net/http"

	"github.com/paubox/paubox-flagr/pkg/config"
)

func getSubjectFromRequest(r *http.Request) string {
	if r == nil {
		return ""
	}

	var token = r.Context().Value(config.TokenContextKey).(*config.DecodedToken)

	if token.Email != "" {
		return token.Email
	}

	return ""
}
