# Start from the latest golang base image
FROM golang:1.16-alpine as build_base


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

FROM alpine:3.14
RUN apk add ca-certificates

COPY --from=build_base /app/main /app/main

RUN chmod +x /app/main

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/main"]