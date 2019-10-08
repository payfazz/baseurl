package baseurl_test

import (
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/payfazz/baseurl"
)

func TestParser1(t *testing.T) {
	testData := []struct {
		expected string
		header   string
	}{
		{"https://example.com/a/b/c", "https://example.com/a/b/c"},
		{"https://example.com/a/b/c", "https://example.com/a/b/c/"},
		{"https://example.com/a/%2fb/c", "https://example.com/a/%2fb/c/"},
		{"", "h      s:"},
	}

	for i := 0; i < len(testData); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expected := testData[i].expected
			header := testData[i].header

			r := httptest.NewRequest("GET", "/a", nil)
			r.Header.Set("X-Base-Url", header)

			found := baseurl.Get(r)
			if found != expected {
				t.Fatalf(`found "%s", expected "%s"`, found, expected)
			}

		})
	}
}

func TestParser2(t *testing.T) {
	r := httptest.NewRequest("GET", "/a", nil)
	r.Header.Set("X-Base-Url", "::::::::invalid value:::::::")

	if baseurl.Get(r) != "" {
		t.Fatalf("should return empty string")
	}
}

func TestParser3(t *testing.T) {
	r := httptest.NewRequest("GET", "/a", nil)
	r.Header.Set("X-Base-Url", "/without/schema/and/host")

	if baseurl.Get(r) != "" {
		t.Fatalf("should return empty string")
	}
}
func TestParser4(t *testing.T) {
	r := httptest.NewRequest("GET", "/a", nil)
	r.Header.Set("X-Base-Url", "http://a:b@c.d/e/f")

	if baseurl.Get(r) != "" {
		t.Fatalf("should return empty string")
	}
}
func TestCurrent(t *testing.T) {
	testData := []struct {
		expected string
		header   string
		path     string
	}{
		{"https://example.com/a/b/c/", "https://example.com/a/b/c", "/"},
		{"https://example.com/a/b/c/d", "https://example.com/a/b/c/", "/d"},
		{"https://example.com/a/%2fb/c/d/", "https://example.com/a/%2fb/c/", "/d/"},
		{"https://example.com/a/%2fb/c/d/%2fe/", "https://example.com/a/%2fb/c/", "/d/%2fe/"},
		{"https://example.com/a/b/c/d/%2fe/", "https://example.com/a/b/c/", "/d/%2fe/"},
		{"", "h      s:", "/"},
	}

	for i := 0; i < len(testData); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			expected := testData[i].expected
			header := testData[i].header
			path := testData[i].path

			r := httptest.NewRequest("GET", path, nil)
			r.Header.Set("X-Base-Url", header)

			found := baseurl.Current(r)
			if found != expected {
				t.Fatalf(`found "%s", expected "%s"`, found, expected)
			}

		})
	}
}