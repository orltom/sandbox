FROM gcr.io/distroless/static:nonroot@sha256:59d91a17dbdd8b785e61da81c9095b78099cad8d7757cc108f49e4fb564ef8b3

ARG TARGETARCH

WORKDIR /

COPY bin/golang-http-example ./golang-http-example

ENTRYPOINT ["/golang-http-example"]

ARG GIT_HASH
ARG BUILD_DATE

LABEL \
	org.opencontainers.image.title="GO webapplication example" \
	org.opencontainers.image.description="GO webapplication example" \
	org.opencontainers.image.documentation="https://github.com/orltom/golang-http-example" \
	org.opencontainers.image.revision="${GIT_HASH}" \
	org.opencontainers.image.created="${BUILD_DATE}"