# Bike Sharing App

This is a simple bike sharing web application powered by React in the frontend and Go in the backend. One could see the location of available bikes and rent/return a bike via the web interface.

## Tech Stack

This app uses a number of third party open-source tools:

### Frontend
- [Vite](https://vitejs.dev/) for building the [React](https://reactjs.org/) frontend.
- [Antd](https://ant.design/) for the UI components.
- [React-leaflet](https://react-leaflet.js.org/) for locating bike in OpenStreetMap.
- [UUID](https://github.com/uuidjs/uuid) for generating using session token.

### Backend
- [Air](https://github.com/cosmtrek/air) for live reloading the [Go](https://go.dev/) backend.
- [Fiber](https://docs.gofiber.io/) as web framework.
- [MongoDB](https://www.mongodb.com/) for data storage.
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/) to work with MongoDB in Go.
- [Go Validator](https://github.com/go-playground/validator) to validate request body.

### Docs
- [Redocly](https://redocly.com/) for generating API documentation.
- [Extract Mongo Schema](https://github.com/perak/extract-mongo-schema) for extracting and generating MongoDB Schema.

## Getting started

### Requirements
You must install [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) to run the application.

### Up and Running
Run following commands from root directory of the project to run the overall application.
```shell
docker-compose build
docker-compose up -d
```

It will build and start the docker image written for the app (which can be seen in [Dockerfile](https://github.com/SRatna/bike-sharing-app/blob/main/Dockerfile)) and also runs MongoDB docker image. We also import some dummy data in MongoDB using `mongoimport`. We have implemented multi-stage builds in the Dockerfile to automate the process of building frontend builds and generating API doc and DB schema.

- The web app can be loaded by visiting [http://localhost:3000/](http://localhost:3000/).
- The API doc can be viewed by visiting [http://localhost:3000/api-doc.html](http://localhost:3000/api-doc.html).
- The MongoDB Schema can be viewed by visiting [http://localhost:3000/schema.html](http://localhost:3000/schema.html).
