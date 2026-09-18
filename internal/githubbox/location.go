// Package githubbox implements the redirects served by githubbox.com.
package githubbox

import "strings"

const prefix = "https://codesandbox.io/s/github/"

// getCodeSandboxLocation returns the CodeSandbox URL for a GitHub path, or "" when the path
// does not point at something CodeSandbox can open.
func getCodeSandboxLocation(path string) string {
	var parts []string
	for _, part := range strings.Split(strings.TrimPrefix(path, "/"), "/") {
		if part != "" {
			parts = append(parts, part)
		}
	}

	switch len(parts) {
	case 0, 1, 3:
		return ""
	case 2:
		return prefix + strings.Join(parts, "/")
	}

	if parts[2] == "tree" {
		return prefix + strings.Join(parts, "/")
	}

	if parts[2] == "blob" {
		return prefix +
			strings.Join([]string{parts[0], parts[1], "tree", parts[3]}, "/") +
			"?file=/" + strings.Join(parts[4:], "/")
	}

	return ""
}
