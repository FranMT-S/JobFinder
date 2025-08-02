package scraper

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
)

func MockServer(fileDir string) (*httptest.Server, string, error) {
	htmlBytes, err := os.ReadFile(fileDir)
	if err != nil {
		return nil, "", err
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(htmlBytes)
	}))

	parsedURL, err := url.Parse(server.URL)
	if err != nil {
		return nil, "", err
	}

	host := parsedURL.Host
	return server, host, nil
}

func DisabledMockServer(server *httptest.Server) {
	server.Close()
}
