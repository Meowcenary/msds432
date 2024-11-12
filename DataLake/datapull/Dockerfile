# Start with the official Golang image to compile the application
FROM golang:1.23 AS build

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application as a static binary
# CGO_ENABLED=0 disables Cgo, which means Go will not use any C libraries, creating a fully static binary.
# GOOS=linux ensures compatibility with Linux when cross-compiling (in this case, optional but good practice).
RUN CGO_ENABLED=0 GOOS=linux go build -o datapull ./cmd

# Use a minimal base image for the runtime
FROM alpine:latest
WORKDIR /root/

# Copy the compiled Go binary from the build stage
COPY --from=build /app/datapull .

# Define the default command to run the service
CMD ["./datapull"]
