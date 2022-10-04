FROM node:18-10-alpine3.15 as build
WORKDIR /app
ENV NODE_ENV="production"
COPY package.json ./
COPY yarn.lock ./
RUN yarn
COPY . ./
RUN yarn build

FROM nginx:1.23-alpine as deploy
WORKDIR /usr/share/nginx/html

RUN rm -rf ./*
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/build .

ENTRYPOINT [ "nginx", "-g", "daemon off;" ]