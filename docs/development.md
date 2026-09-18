## githubbox

The site is a static [Vercel](https://vercel.com/) deployment which redirects
visitors to CodeSandbox. There is no application code: the redirects are
declared in `vercel.json` and handled by Vercel's CDN, so they run in every
region without invoking a function.

See [redirect.md](./redirect.md) for the redirect setup.

## Layout

- `vercel.json` holds the redirect rules
- `public/robots.txt` is served as a static file. Vercel serves this project's
  static files out of `public/`, so a `robots.txt` at the repo root is not
  reachable — it has to live here

## Editing

To edit the redirect logic, open `vercel.json`. The rules are matched in order,
first match wins:

| Source                          | Goes to                                    |
| ------------------------------- | ------------------------------------------ |
| `/` and `/index.html`           | this repo on GitHub                        |
| `/:owner/:repo/blob/:ref/:file+` | the sandbox for `:ref`, with `:file` opened |
| `/:owner/:repo/tree/:rest+`     | the sandbox for that branch, tag or commit |
| `/:owner/:repo`                 | the sandbox for the repo's default branch  |

Anything else falls through to a 404, which is what we want for paths like
`/issues` or `/vercel/ms/stargazers`.

The `+` modifier requires at least one segment, so `/vercel/ms/tree` alone is a
404 rather than a redirect to an empty ref.

## Running locally

```
vercel dev
```

Then check http://localhost:3000/vercel/ms

## Testing

There is nothing to unit test, so verify against a deployment instead. Point
`BASE` at a preview URL and check the `location` headers:

```
BASE=https://githubbox.com
for path in \
  / \
  /vercel/ms \
  /vercel/ms/ \
  /vercel/ms/tree/main \
  /vercel/ms/tree/2.1.1 \
  /vercel/ms/blob/main/package.json \
  /vercel/ms/blob/main/lib/index.js \
  /robots.txt \
  /issues \
  /vercel/ms/stargazers
do
  echo "$path -> $(curl -sI "$BASE$path" | head -n1 | tr -d '\r')"
done
```

The last two should be 404s. `/robots.txt` should be a 200.

## Deploying

### Preview

```
vercel
```

Then check the preview URL that gets printed.

### Production

```
vercel --prod
```

Then check https://githubbox.com

## Test URLs

All these forms of URLs are supported.

- main [githubbox.com/vercel/ms](https://githubbox.com/vercel/ms)
- branch [githubbox.com/vercel/ms/tree/main](https://githubbox.com/vercel/ms/tree/main)
- tree [githubbox.com/vercel/ms/tree/2.1.1](https://githubbox.com/vercel/ms/tree/2.1.1)
- commit [githubbox.com/vercel/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf](https://githubbox.com/vercel/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf)
- branch with file [githubbox.com/vercel/ms/blob/main/package.json](https://githubbox.com/vercel/ms/blob/main/package.json)
- tree with file [githubbox.com/vercel/ms/blob/2.0.0/index.js](https://githubbox.com/vercel/ms/blob/2.0.0/index.js)
- commit with file [githubbox.com/vercel/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js](https://githubbox.com/vercel/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js)
