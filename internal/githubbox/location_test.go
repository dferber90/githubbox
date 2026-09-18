package githubbox

import "testing"

func TestGetCodeSandboxLocation(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"root", "/", ""},
		{"repo root", "/zeit/ms", "https://codesandbox.io/s/github/zeit/ms"},
		{"repo root trailing slash", "/zeit/ms/", "https://codesandbox.io/s/github/zeit/ms"},
		{"repo branch", "/zeit/ms/tree/master", "https://codesandbox.io/s/github/zeit/ms/tree/master"},
		{"repo branch trailing slash", "/zeit/ms/tree/master/", "https://codesandbox.io/s/github/zeit/ms/tree/master"},
		{"repo tree", "/zeit/ms/tree/2.1.1", "https://codesandbox.io/s/github/zeit/ms/tree/2.1.1"},
		{"repo tree trailing slash", "/zeit/ms/tree/2.1.1/", "https://codesandbox.io/s/github/zeit/ms/tree/2.1.1"},
		{
			"repo commit",
			"/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf",
			"https://codesandbox.io/s/github/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf",
		},
		{
			"repo commit trailing slash",
			"/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf/",
			"https://codesandbox.io/s/github/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf",
		},
		{
			"branch with file",
			"/zeit/ms/blob/master/package.json",
			"https://codesandbox.io/s/github/zeit/ms/tree/master?file=/package.json",
		},
		{
			"tree with file",
			"/zeit/ms/blob/2.0.0/index.js",
			"https://codesandbox.io/s/github/zeit/ms/tree/2.0.0?file=/index.js",
		},
		{
			"commit with file",
			"/zeit/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js",
			"https://codesandbox.io/s/github/zeit/ms/tree/adf1eb282d29fe3c405d205a3854177b86a97c1f?file=/index.js",
		},
		{"page which is not a repo", "/issues", ""},
		{"page which is not a repo, trailing slash", "/issues/", ""},
		{"repo page which is not a repo", "/zeit/ms/stargazers", ""},
		{"repo page which is not a repo, trailing slash", "/zeit/ms/stargazers/", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCodeSandboxLocation(tt.path); got != tt.want {
				t.Errorf("getCodeSandboxLocation(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
