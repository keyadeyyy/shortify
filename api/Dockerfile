# -------- STAGE 1: Build the Go application --------
# Use an official minimal Go image with Alpine Linux for building
#To compile your Go app into a standalone binary

FROM golang:alpine AS builder

# Create a working directory for your Go source code
RUN mkdir /build

# Copy all local files (source, go.mod, etc.) to /build inside the container
ADD . /build/

# Set the working directory for the build process
WORKDIR /build

# Build the Go binary. -o main names the output binary "main"
RUN go build -o main .



# -------- STAGE 2: Run the built binary in a clean image --------
#To run the already-built app safely & efficiently
FROM alpine

# Create a new non-root user named 'appuser' with home directory /app
RUN adduser -S -D -H -h /app appuser

# Switch to the non-root user for better security
USER appuser


# Copy the compiled binary from the 'builder' stage
COPY --from=builder /build/main /app/

# Set the working directory inside the container to /app
WORKDIR /app

# Expose port 3000 to the host (match your app's Fiber.Listen port)
EXPOSE 3000

# Define the command to run your Go app
CMD ["./main"]
