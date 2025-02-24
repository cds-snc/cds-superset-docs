FROM oven/bun:1.2.2-slim AS install

RUN mkdir -p /tmp/build
WORKDIR /tmp/build

COPY . .
RUN bun install --frozen-lockfile --production
RUN bun build --entrypoints ./index.ts --outdir ./dist --target=bun

FROM oven/bun:1.2.2-distroless AS release
WORKDIR /var/task

COPY --from=install /tmp/build/dist index.js
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.9.0 /lambda-adapter /opt/extensions/lambda-adapter
COPY ./public ./public
COPY ./views ./views

ENV PORT=5000
ENTRYPOINT [ "bun", "run", "index.js" ]