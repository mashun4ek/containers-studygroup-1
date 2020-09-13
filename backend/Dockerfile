# build stage
FROM golang:1.15.0-alpine3.12 AS builder
# Here you can install dependencies, for example let's install git
RUN apk update && apk --no-cache add git
# Set the current working directory inside the container 
WORKDIR /src
# Copy go mod and sum files 
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 
# Copy the source from the current directory to the working Directory inside the container 
COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o goapp .

# final stage
FROM alpine:3.12.0
RUN apk --no-cache add ca-certificates
WORKDIR /app
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /src/goapp .  
# Expose port 8080 to the outside world
EXPOSE 8080
#Command to run the executable
ENTRYPOINT ["./goapp"]

# 1. Build the image
# docker build -t masha:2.0 .
# 2. Run the container, but don't forget to pass env variables from .env-db file
# docker run -it --env-file=.env-db --rm -p 8080:8080 masha:2.0