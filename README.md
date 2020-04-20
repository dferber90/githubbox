## githubbox

## Running locally

```
yarn dev
```

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
