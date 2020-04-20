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
