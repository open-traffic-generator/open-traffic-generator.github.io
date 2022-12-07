# Open Traffic Generator APIs & Data Models

This repo contains the content for the [Open Traffic Generator APIs and Data Models](https://open-traffic-generator.github.io/) web-site.  It is built using the [Material](https://squidfunk.github.io/mkdocs-material/getting-started/) theme for [MkDocs](https://www.mkdocs.org/).

Update contents in the `docs` directory and verify locally prior to pushing to main branch of this repo on GitHub.  Site will automatically update.

To verify contents locally, make sure that you have Material [installed](https://squidfunk.github.io/mkdocs-material/getting-started/) and then run the following command :

```sh
$ mkdocs build
```

You can point your browser to `index.html` in the `site` directory to view it.

## Submodules

Parts of the `docs` hierarchy are coming from submodules. To update content of the submodules to the most recent one, use:

```Shell
git submodule update --remote
```