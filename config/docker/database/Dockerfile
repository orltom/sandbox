FROM postgres:alpine

ADD build/jokes.sql /docker-entrypoint-initdb.d

RUN chmod a+r /docker-entrypoint-initdb.d/*

ARG GIT_HASH
ARG BUILD_DATE

LABEL \
	org.opencontainers.image.title="GO webapplication example" \
	org.opencontainers.image.description="GO webapplication example" \
	org.opencontainers.image.documentation="https://github.com/orltom/golang-http-example" \
	org.opencontainers.image.revision="${GIT_HASH}" \
	org.opencontainers.image.created="${BUILD_DATE}"