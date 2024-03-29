FROM playcourt/golang:1.21 AS builder
WORKDIR /usr/src/app
USER root
RUN apk add --update --no-cache build-base
COPY go.mod .
COPY go.sum .
COPY Makefile .
RUN make install
COPY . .
RUN make build

FROM playcourt/alpine:base
WORKDIR /usr/src/app
USER user
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG METABASE_SECRET_KEY
ARG METABASE_SITE_URL

RUN touch .env
RUN echo "DB_HOST=${DB_HOST}" >> .env
RUN echo "DB_PORT=${DB_PORT}" >> .env
RUN echo "DB_USER=${DB_USER}" >> .env
RUN echo "DB_PASSWORD=${DB_PASSWORD}" >> .env
RUN echo "DB_NAME=${DB_NAME}" >> .env
RUN echo "METABASE_SECRET_KEY=${METABASE_SECRET_KEY}" >> .env
RUN echo "METABASE_SITE_URL=${METABASE_SITE_URL}" >> .env

COPY --from=builder --chown=user:root /usr/src/app/main /usr/src/app
CMD ["./main"]