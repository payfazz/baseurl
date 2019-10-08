package baseurl

import (
	"net/http"
	"net/url"
)

// Get request and return base url.
// return empty string if base url cannot be infered from request header.
// base url will not have trailing slash
func Get(r *http.Request) string {
	header := r.Header.Get("X-Base-Url")
	if header == "" {
		return ""
	}

	// Must valid url
	base, err := url.Parse(header)
	if err != nil {
		return ""
	}

	// must have schema and host
	if base.Scheme == "" || base.Host == "" {
		return ""
	}

	// must not containn extra field
	if base.Opaque != "" || base.User != nil || base.ForceQuery || base.RawQuery != "" || base.Fragment != "" {
		return ""
	}

	ret := base.String()

	// remove trailing slash
	if len := len(ret); len > 0 && ret[len-1] == '/' {
		ret = ret[:len-1]
	}

	return ret
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
