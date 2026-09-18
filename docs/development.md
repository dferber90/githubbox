## githubbox

The site is a Go application deployed on [Vercel](https://vercel.com/) which
redirects visitors to CodeSandbox.

See [redirect.md](./redirect.md) for the redirect setup.

## Layout

- `cmd/server/main.go` is the entrypoint Vercel's Go framework preset builds
  and runs, listening on `$PORT`
- `internal/githubbox` holds the actual logic and its tests
- `vercel.json` sets the `go` framework preset

## Editing

To edit the redirect logic, open `internal/githubbox/location.go`.

## Testing

```
go test ./...
```

## Running locally

```
go run ./cmd/server
```

Then check http://localhost:3000/zeit/ms

It listens on `$PORT`, defaulting to 3000, so `PORT=8080 go run ./cmd/server`
works too. To run it the way Vercel does, use the
[Vercel CLI](https://vercel.com/docs/cli) instead:

```
vercel dev
```

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

- main [githubbox.com/zeit/ms](https://githubbox.com/zeit/ms)
- branch [githubbox.com/zeit/ms/tree/main](https://githubbox.com/zeit/ms/tree/main)
- tree [githubbox.com/zeit/ms/tree/2.1.1](https://githubbox.com/zeit/ms/tree/2.1.1)
- commit [githubbox.com/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf](https://githubbox.com/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf)
- branch with file [githubbox.com/zeit/ms/blob/main/package.json](https://githubbox.com/zeit/ms/blob/main/package.json)
- tree with file [githubbox.com/zeit/ms/blob/2.0.0/index.js](https://githubbox.com/zeit/ms/blob/2.0.0/index.js)
- commit with file [githubbox.com/zeit/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js](https://githubbox.com/zeit/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js)
