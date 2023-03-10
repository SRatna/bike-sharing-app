# frontend build
FROM node:18-alpine AS frontendBuild

WORKDIR /app

COPY frontend/package.json ./
COPY frontend/package-lock.json ./

RUN npm install

COPY ./frontend ./

RUN npm run build

# api doc build
FROM node:18-alpine AS apiDocBuild

WORKDIR /app

COPY docs/package.json ./
COPY docs/package-lock.json ./

RUN npm install

COPY ./docs/openapi.yaml ./openapi.yaml

RUN npm run build

# build mongo schema
FROM node:18-alpine AS mongoSchemaBuild

WORKDIR /app

COPY docs/package.json ./
COPY docs/package-lock.json ./

RUN npm install

COPY ./docs/schema.json ./schema.json

RUN mkdir dist 

RUN npm run build-schema

# backend build and run
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./db ./db
COPY ./handlers ./handlers

RUN go build -o /bike-app

COPY --from=frontendBuild /app/dist ./dist
COPY --from=apiDocBuild /app/dist/api-doc.html ./dist/api-doc.html
COPY --from=mongoSchemaBuild /app/dist/schema.html ./dist/schema.html

EXPOSE 3000

CMD [ "/bike-app" ]