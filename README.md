# CDS Superset Docs

Documentation site for [CDS Superset](https://superset.cds-snc.ca/). Content is managed in WordPress and retrieved via its API.

The app is built using [Express.js](https://expressjs.com/) and uses the [GC Design System](https://design-system.alpha.canada.ca/) for its frontend. Hosting is done using a Lambda function with content caching provided by CloudFront.

## Running locally
```sh
# Install bun or use the devcontainer
# https://bun.sh/docs/installation
cd app
cp .env.example .env # and set your values
bun install
bun start
```
