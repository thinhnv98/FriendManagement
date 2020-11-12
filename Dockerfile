# Use the official image as a parent image.
FROM golang:alpine as builder

ENV GO111MODULE=on

# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Thinh.NguyenV <thinh.nguyenv@s3corp.com.vn>"

# Set the working directory.
WORKDIR /app

# Copy the file from your host to your current location.
COPY .env .
COPY go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env ./.env

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]