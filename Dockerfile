FROM golang:latest as builder
LABEL maintainer="leonardo.siagian@koinworks.com"

ENV GO111MODULE=on

ENV GOPRIVATE=github.com/koinworks

ARG GITHUB_USERNAME
ARG GITHUB_ACCESS_TOKEN

RUN echo "machine github.com login $GITHUB_USERNAME password $GITHUB_ACCESS_TOKEN" >  ~/.netrc

ENV APP bivrost-task2

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE ${PORT}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/${APP} main.go
ENTRYPOINT /out/${APP}