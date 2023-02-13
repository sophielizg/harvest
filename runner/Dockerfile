# Builder
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git bash

ENV USER=app
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /usr/src/app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
RUN go mod verify

COPY . .
ENV GOOS=linux 
ENV GOARCH=amd64 
ENV CGO_ENABLED=0
RUN go build -o runner cmd/harvest/main.go


# FINAL
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /usr/src/app /

USER app:app

ENTRYPOINT ["/runner"]
