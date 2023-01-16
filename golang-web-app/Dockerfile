# Specifies a parent image
FROM golang:1.19.2-alpine3.16 AS builder

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY go.mod .
COPY go.sum .

ENV GOPROXY https://goproxy.cn

# Installs Go dependencies
RUN go mod download

COPY . .

# Builds your app with optional configuration
RUN go build -o golang-web-app

FROM alpine:3.16

WORKDIR /

COPY --from=builder /app .

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
CMD [ "./golang-web-app" ]