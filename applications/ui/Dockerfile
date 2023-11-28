FROM node:21 AS build-env
COPY package.json package-lock.json* ./
RUN npm ci
RUN npm cache clean --force
COPY . /.
RUN npm run build

FROM nginxinc/nginx-unprivileged:1.25.3-alpine-slim
COPY --from=build-env --chown=nginx:nginx build/ /usr/share/nginx/html
COPY nginx.conf /etc/nginx/templates/default.conf.template

