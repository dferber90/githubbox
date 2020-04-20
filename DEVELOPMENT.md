## githubbox

The site uses Cloudflare Workers to redirect visitors to CodeSandbox.

See [redirect.md](./redirect.md) for the redirect setup.

## Running locally

```
yarn start
```

Then open `http://localhost:5000`.

> You'll probably only be redirected to github.com anyways, since that is all `index.html' does.

## Editing the worker

To edit the worker, open `workers-site/index.js`.
The `workers-site` directory contains the Cloudflare worker.

## Building and publishing a new version

You need to install [wrangler](https://github.com/cloudflare/wrangler) globally.

### Publishing to development

```
yarn pub
```

Then check https://githubbox.dferber.workers.dev/

### Publishing to production

```
yarn pub:prod
```

Then check https://githubbox.com

## Test URLs

All these forms of URLs are supported.

- main [githubbox.com/zeit/ms](https://githubbox.com/zeit/ms)
- branch [githubbox.com/zeit/ms/tree/master](https://githubbox.com/zeit/ms/tree/master)
- tree [githubbox.com/zeit/ms/tree/2.1.1](https://githubbox.com/zeit/ms/tree/2.1.1)
- commit [githubbox.com/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf](https://githubbox.com/zeit/ms/tree/7920885eb232fbe7a5efdab956d3e7c507c92ddf)
- branch with file [githubbox.com/zeit/ms/blob/master/package.json](https://githubbox.com/zeit/ms/blob/master/package.json)
- tree with file [githubbox.com/zeit/ms/blob/2.0.0/index.js](https://githubbox.com/zeit/ms/blob/2.0.0/index.js)
- commit with file [githubbox.com/zeit/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js](https://githubbox.com/zeit/ms/blob/adf1eb282d29fe3c405d205a3854177b86a97c1f/index.js)
