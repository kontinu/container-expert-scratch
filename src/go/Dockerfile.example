FROM golang:1.14.3 AS build

WORKDIR /compose/main

COPY . ./
COPY main.go main.go
RUN CGO_ENABLED=0 go build -o backend main.go

FROM alpine:3
COPY --from=build /compose/main/backend /usr/local/bin/backend
ENTRYPOINT  ["/usr/local/bin/backend"]
