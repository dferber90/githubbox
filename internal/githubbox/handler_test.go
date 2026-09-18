package githubbox

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func request(path string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	Handle(w, r)
	return w
}

func TestHandleRedirects(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		location string
	}{
		{"root", "/", "https://github.com/dferber90/githubbox"},
		{"index.html", "/index.html", "https://github.com/dferber90/githubbox"},
		{"repo", "/zeit/ms", "https://codesandbox.io/s/github/zeit/ms"},
		{"branch", "/zeit/ms/tree/master", "https://codesandbox.io/s/github/zeit/ms/tree/master"},
		{
			"blob",
			"/zeit/ms/blob/master/package.json",
			"https://codesandbox.io/s/github/zeit/ms/tree/master?file=/package.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := request(tt.path)
			if w.Code != http.StatusFound {
				t.Errorf("status = %d, want %d", w.Code, http.StatusFound)
			}
			if got := w.Header().Get("Location"); got != tt.location {
				t.Errorf("Location = %q, want %q", got, tt.location)
			}
		})
	}
}

func TestHandleRobotsTxt(t *testing.T) {
	w := request("/robots.txt")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}
	want := "# https://www.robotstxt.org/robotstxt.html\nUser-agent: *\nDisallow:"
	if got := w.Body.String(); got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestHandleNotFound(t *testing.T) {
	for _, path := range []string{"/issues", "/zeit/ms/stargazers"} {
		w := request(path)
		if w.Code != http.StatusNotFound {
			t.Errorf("status for %q = %d, want %d", path, w.Code, http.StatusNotFound)
		}
	}
}
