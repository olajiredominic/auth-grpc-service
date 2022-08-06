# syntax=docker/dockerfile:1
FROM golang:1.18-bullseye
RUN go install github.com/beego/bee/v2@latest

# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/app
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
COPY go.mod /go/src/app
COPY go.sum /go/src/app
RUN go mod download

COPY ./ /go/src/app
WORKDIR "$APP_HOME/cmd"

EXPOSE 5001
CMD ["bee", "run"]
# ENTRYPOINT "bee run"