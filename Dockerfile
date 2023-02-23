# frontend build
FROM node:18-alpine AS frontendBuild

WORKDIR /app

COPY frontend/package.json ./
COPY frontend/package-lock.json ./

RUN npm install

COPY ./frontend ./

RUN npm run build

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

EXPOSE 3000

CMD [ "/bike-app" ]