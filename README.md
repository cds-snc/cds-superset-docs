# CDS Superset Docs

Documentation site for [CDS Superset](https://superset.cds-snc.ca/). Content is managed in WordPress and retrieved via its API.

The app is built using [Express.js](https://expressjs.com/) and uses the [GC Design System](https://design-system.alpha.canada.ca/) for its frontend. Hosting is done using a [Lambda function](./terragrunt/aws/lambda.tf) with content caching provided by a [CloudFront distribution](./terragrunt/aws/cloudfront.tf).

## Running locally

1. [Install Go](https://go.dev/doc/install) or use the devcontainer.
1. Update the values in `./app/.env.example` and set them your terminal environment.

```sh
cd app
go mod download
go run cmd/server/main.go
```
