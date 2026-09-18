package githubbox

import (
	"net/http"
	"strings"
)

const repoURL = "https://github.com/dferber90/githubbox"

var robotsTxt = strings.Join([]string{
	"# https://www.robotstxt.org/robotstxt.html",
	"User-agent: *",
	"Disallow:",
}, "\n")

// Handle serves a single request: it redirects GitHub paths to their
// CodeSandbox equivalent and answers the few static paths the site has.
func Handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" || path == "/index.html" {
		http.Redirect(w, r, repoURL, http.StatusFound)
		return
	}

	if path == "/robots.txt" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(robotsTxt))
		return
	}

	if location := getCodeSandboxLocation(path); location != "" {
		http.Redirect(w, r, location, http.StatusFound)
		return
	}

	http.Error(w, "Not found", http.StatusNotFound)
}
