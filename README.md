Quickly open any GitHub repo in CodeSandbox

1. Go to a repo in GitHub
2. Replace `github.com` with `githubbox.com`
3. There's no step three

<p align="center">
  <img src="./docs/demo-cra.gif" width="718">
</p>

## What

Suppose you're on `github.com` looking at a repo.

Now you want to explore the code of that repo in a [CodeSandbox](https://codesandbox.io/).

Just add **box** to your url and the repo will open there.

## Try

As you're on GitHub at the moment, you can add `box` to the current URL.

```diff
-https://github.com/dferber90/githubbox
+https://githubbox.com/dferber90/githubbox
```

You'll then see this README on CodeSandbox.

_This also works when you're on a branch, a tag or on a file's blob._

## Lazy

If you don't want to type, you can click on this link instead:

[githubbox.com/dferber90/githubbox](https://githubbox.com/dferber90/githubbox)

## Badge

You can add an _"**Open in CodeSandbox**"_ badge to your README if you own a repo.

This allows anybody to open your repo on CodeSandbox with one click.

#### Example

[![Open in CodeSandbox](https://img.shields.io/badge/Open%20in-CodeSandbox-blue?style=flat-square&logo=codesandbox)](https://githubbox.com/dferber90/githubbox)

#### Markdown

```markdown
[![Open in CodeSandbox](https://img.shields.io/badge/Open%20in-CodeSandbox-blue?style=flat-square&logo=codesandbox)](https://githubbox.com/dferber90/githubbox)
```

Replace `https://githubbox.com/dferber90/githubbox` with your own githubbox.com URL.

See [shields.io](https://shields.io/) for badge style customization options.

## Development

See [DEVELOPMENT.md](./docs/development.md).

## Who

[@dferber90](https://twitter.com/dferber90)
