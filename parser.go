package baseurl

import (
	"net/http"
	"net/url"
)

// Parse request and return base url.
// return nil if base url cannot be infered from request header.
func Parse(r *http.Request) *url.URL {
	baseFromHeader := r.Header.Get("X-Base-Url")
	if baseFromHeader == "" {
		return nil
	}

	baseURL, err := url.Parse(baseFromHeader)
	if err != nil {
		return nil
	}

	// X-Base-Url must have schema and host
	if baseURL.Scheme == "" || baseURL.Host == "" {
		return nil
	}

	// remove trailing slash
	escapedPath := baseURL.EscapedPath()
	if escapedPathLen := len(escapedPath); escapedPathLen > 0 && escapedPath[escapedPathLen-1] == '/' {
		baseURL.Path = baseURL.Path[:len(baseURL.Path)-1]
		if escapedPath == baseURL.RawPath {
			baseURL.RawPath = baseURL.RawPath[:len(baseURL.RawPath)-1]
		} else {
			baseURL.RawPath = ""
		}
	}

	return baseURL
}

// Current return current url.Current
// return nil if Parse return nil.
func Current(r *http.Request) *url.URL {
	base := Parse(r)
	if base == nil {
		return nil
	}

	current, err := url.Parse(base.String() + r.URL.EscapedPath())
	if err != nil {
		return nil
	}

	return current
}
