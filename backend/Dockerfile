# Use Golang image as the base image

FROM golang:1.24 AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go mod tidy
RUN go build -o main .

# Check that the binary exists in the build image (for debugging)
RUN ls -l /app

# Use a smaller base image for the final container
FROM ubuntu:22.04

# Set the working directory inside the final container
WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=build /app/main .

# Check that the binary exists in the final image (for debugging)
RUN ls -l /root/

# Expose port 7000 for the Go application
EXPOSE 7000

# Command to run the Go application
CMD ["./main"]
