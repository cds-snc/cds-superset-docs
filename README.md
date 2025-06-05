# CDS Superset Docs

Documentation site for [CDS Superset](https://superset.cds-snc.ca/). Content is managed in WordPress and retrieved via its API.

The app is built using [Go](https://go.dev/) and uses the [GC Design System](https://design-system.alpha.canada.ca/) for its frontend. Hosting is done using a [Lambda function](./terragrunt/aws/lambda.tf) with content caching provided by a [CloudFront distribution](./terragrunt/aws/cloudfront.tf).

## Running locally

To run the project locally [install Go](https://go.dev/doc/install) or use the [devcontainer](https://containers.dev/supporting).  Then, from a terminal:

```sh
cd app
cp .env.example .env # and set your values
make run
```
