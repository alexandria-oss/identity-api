# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer Info
LABEL maintainer="Alonso R <luis.alonso.16@hotmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/alexandria-oss/alexandria/identity-api/

# Copy go mod files
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o identity ./cmd/prod/main.go


######## Start a new stage from scratch #######
FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/alexandria-oss/alexandria/identity-api .
COPY --from=builder /go/src/github.com/alexandria-oss/alexandria/identity-api/config/alexandria-config.yaml .

CMD ["mkdir .aws"]
COPY --from=builder /go/src/github.com/alexandria-oss/alexandria/identity-api/config/.aws ./.aws

# Expose port 8080 -> HTTP & 9090 -> gRPC to the outside world
EXPOSE 8080
EXPOSE 9090

# Command to run the executable
CMD ["./identity"]
