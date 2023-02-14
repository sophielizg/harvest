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

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.9

COPY api/go.mod .
COPY api/go.sum .
RUN go mod download
RUN go mod verify

COPY api .
ENV GOOS=linux 
ENV GOARCH=amd64 
ENV CGO_ENABLED=0
RUN swag init --parseDependency -g ./cmd/harvest/main.go
RUN go build -o server cmd/harvest/main.go


# FINAL
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /usr/src/app/server /server
COPY config /config

USER app:app

ENV PORT=8080

CMD ["/server"]
