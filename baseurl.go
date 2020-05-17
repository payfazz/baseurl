package baseurl

import (
	"net/http"
	"net/url"
)

// Get request and return base url.
// return empty string if base url cannot be inferred from request header.
// base url will not have trailing slash
func Get(r *http.Request) string {
	header := r.Header.Get("X-Base-Url")
	if header == "" {
		return ""
	}

	// Must valid url
	url, err := url.Parse(header)
	if err != nil {
		return ""
	}

	// must have schema and host
	if url.Scheme == "" || url.Host == "" {
		return ""
	}

	// must not containn extra field
	if url.Opaque != "" || url.User != nil || url.ForceQuery || url.RawQuery != "" || url.Fragment != "" {
		return ""
	}

	// remove trailing slash
	if header[len(header)-1] == '/' {
		header = header[:len(header)-1]
	}

	return header
}

// MustGet is like Get, but instead return empty string
// it will recreate url from r if Get return empty string
func MustGet(r *http.Request) string {
	if base := Get(r); base != "" {
		return base
	}

	schema := "http"
	if r.TLS != nil {
		schema = "https"
	}

	host := r.Host
	if host == "" {
		host = "localhost"
	}

	return schema + "://" + host
}

// Current return current url
// return empty string if Get return empty string.
func Current(r *http.Request) string {
	base := Get(r)
	if base == "" {
		return ""
	}

	current, err := url.Parse(base + r.URL.EscapedPath())
	if err != nil {
		return ""
	}

	return current.String()
}
