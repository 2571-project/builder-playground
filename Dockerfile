# Use a lightweight Alpine Linux base image
FROM alpine:latest

# Install necessary dependencies
RUN apk add --no-cache ca-certificates

# Set environment variables with default values
ENV API_LISTEN_ADDR=127.0.0.1 \
    API_LISTEN_PORT=5555 \
    API_SECRET_KEY=5eae315483f028b5cdd5d1090ff0c7618b18737ea9bf3c35047189db22835c48 \
    BEACON_CLIENT_ADDR=http://localhost:4500

# Create a directory for the application
WORKDIR /app

# Copy the builder-playground binary into the container
COPY builder-playground /app/builder-playground

# Ensure the binary is executable
RUN chmod +x /app/builder-playground

# Expose the API port
EXPOSE 5555

# Define the entrypoint to run the application
CMD ["./builder-playground"]
