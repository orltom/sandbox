FROM nginxinc/nginx-unprivileged:1.25.3-alpine-slim@sha256:814b97900dbf3c724e1718789c2bfe4c7caa1c9eddc1caa7806dea0dfdad89c2

COPY build /usr/share/nginx/html

COPY nginx.conf /etc/nginx/templates/default.conf.template

LABEL \
	org.opencontainers.image.title="Sandbox UI" \
	org.opencontainers.image.description="Webapplication to visualize random jokes" \
	org.opencontainers.image.documentation="https://github.com/orltom/sandbox"