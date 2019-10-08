package baseurl_test

import (
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/payfazz/baseurl"
)

func TestParser1(t *testing.T) {
	testData := []struct {
		oke      bool
		expected string
		header   string
	}{
		{true, "https://example.com/a/b/c", "https://example.com/a/b/c"},
		{true, "https://example.com/a/b/c", "https://example.com/a/b/c/"},
		{true, "https://example.com/a/%2fb/c", "https://example.com/a/%2fb/c/"},
		{false, "", "h      s:"},
	}

	for i := 0; i < len(testData); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			oke := testData[i].oke
			expected := testData[i].expected
			header := testData[i].header

			r := httptest.NewRequest("GET", "/a", nil)
			r.Header.Set("X-Base-Url", header)

			url := baseurl.Parse(r)
			if oke {
				found := url.String()
				if found != expected {
					t.Fatalf(`found "%s", expected "%s"`, found, expected)
				}
			} else {
				if url != nil {
					found := url.String()
					t.Fatalf(`found "%s", expected nil`, found)
				}
			}

		})
	}
}

func TestParser2(t *testing.T) {
	r := httptest.NewRequest("GET", "/a", nil)
	r.Header.Set("X-Base-Url", "::::::::invalid value:::::::")

	if baseurl.Parse(r) != nil {
		t.Fatalf("should return nil")
	}
}

func TestParser3(t *testing.T) {
	r := httptest.NewRequest("GET", "/a", nil)
	r.Header.Set("X-Base-Url", "/without/schema/and/host")

	if baseurl.Parse(r) != nil {
		t.Fatalf("should return nil")
	}
}
func TestCurrent(t *testing.T) {
	testData := []struct {
		oke      bool
		expected string
		header   string
		path     string
	}{
		{true, "https://example.com/a/b/c/", "https://example.com/a/b/c", "/"},
		{true, "https://example.com/a/b/c/d", "https://example.com/a/b/c/", "/d"},
		{true, "https://example.com/a/%2fb/c/d/", "https://example.com/a/%2fb/c/", "/d/"},
		{true, "https://example.com/a/%2fb/c/d/%2fe/", "https://example.com/a/%2fb/c/", "/d/%2fe/"},
		{true, "https://example.com/a/b/c/d/%2fe/", "https://example.com/a/b/c/", "/d/%2fe/"},
		{false, "", "h      s:", "/"},
	}

	for i := 0; i < len(testData); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			oke := testData[i].oke
			expected := testData[i].expected
			header := testData[i].header
			path := testData[i].path

			r := httptest.NewRequest("GET", path, nil)
			r.Header.Set("X-Base-Url", header)

			url := baseurl.Current(r)
			if oke {
				found := url.String()
				if found != expected {
					t.Fatalf(`found "%s", expected "%s"`, found, expected)
				}
			} else {
				if url != nil {
					found := url.String()
					t.Fatalf(`found "%s", expected nil`, found)
				}
			}

		})
	}
}
