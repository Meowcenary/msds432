# Start with the official Golang image to compile the application
FROM golang:1.23 AS build

# Set the working directory in the container
WORKDIR /app

# Copy dependency files
COPY ./api/go.mod ./api/go.sum ./

# Download all dependencies
RUN go mod download

# Copy application source code
COPY ./api/ /app/.

# Copy shared code
COPY ./shared/models /app/internal/models/.
COPY ./shared/dbconnector /app/internal/dbconnector/.

# Build the application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api/api ./cmd

# Use a minimal base image for the runtime
FROM alpine:latest
WORKDIR /root/

# Copy the compiled Go binary from the build stage
COPY --from=build /app/api/api /root/

# Define the default command to run the service
CMD ["./api"]
