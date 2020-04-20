## githubbox

## Running locally

```
yarn start
```

## Editing the worker

To edit the worker, open `workers-site/index.js`.
The `workers-site` directory contains the Cloudflare worker.

## Building and publishing a new version

You need to install [wrangler](https://github.com/cloudflare/wrangler) globally.

### Publishing to development

```
yarn build
yarn pub
```

Then check https://githubbox.dferber.workers.dev/

### Publishing to production

```
yarn build
yarn pub:prod
```

Then check https://githubbox.com
