package awair

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

const (
	baseUrlPath = "/v1"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	rootHandler := http.NewServeMux()
	rootHandler.Handle(baseUrlPath+"/", http.StripPrefix(baseUrlPath, mux))
	rootHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("unregistered URL: %s", r.URL.String()), http.StatusInternalServerError)
	})

	server := httptest.NewServer(rootHandler)

	client = NewClient(nil)
	baseUrl, _ := url.Parse(server.URL + baseUrlPath + "/")
	client.baseUrl = baseUrl

	return client, mux, server.Close
}

func stringPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}
