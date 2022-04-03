# Build stage
FROM golang:1.18-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o simplebank main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/simplebank .
COPY --from=build /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD ["/app/simplebank"]

# entrypoint of the docker image with the CMD passed as an argument
# basically it looks like this in the background ["/app/start.sh" "/app/main"] 
# which is then passed as a variable to the exec command in the start.sh script
ENTRYPOINT [ "/app/start.sh"]
