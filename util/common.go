package util

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/rajatxs/go-cosmic-echoes/types"
)

// Sends standard API response
func SendResponse(w *http.ResponseWriter, status int, message string, result interface{}) error {
	payload := &types.StdResponse{}

	payload.StatusCode = status
	payload.Message = message

	if result == nil {
		payload.Result = new(struct{})
	} else {
		payload.Result = result
	}

	(*w).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*w).WriteHeader(status)
	return json.NewEncoder(*w).Encode(payload)
}

func GetAbsoluteUrl(baseUrl string, endpoint string) string {
	var (
		parsedEndpoint *url.URL
		parsedBaseUrl  *url.URL
		err            error
	)

	if len(endpoint) == 0 {
		return ""
	}

	parsedEndpoint, err = url.Parse(endpoint)

	if err != nil {
		return ""
	}

	if len(parsedEndpoint.Scheme) == 0 {
		if parsedBaseUrl, err = url.Parse(baseUrl); err == nil {
			return parsedBaseUrl.ResolveReference(&url.URL{Path: endpoint}).String()
		} else {
			return ""
		}
	} else {
		return endpoint
	}
}
