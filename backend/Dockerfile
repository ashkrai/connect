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

# END

# RUN : docker image prune : After building your image, use this command to remove unused intermediate images eg : <none>
# RUN :  docker run -p 7000:7000 <name>:<version>
#  eg : docker build -t ashkrai/backend:0.4 .       docker run -p 7000:7000 ashkrai/backend:0.4

# NOTES
# In the Dockerfile you provided, you're using a multi-stage build, where two images are involved:
# Stage1 : The build image (golang:1.24): This is used to compile your Go code.
# Stage2 : The final image (ubuntu:22.04): This is used to package and run your Go binary in a smaller, production-ready image.
# The <none> image is an intermediate image created during a multi-stage build, used only for compiling your Go app.
#     It doesn't affect the final image and can be safely deleted.


