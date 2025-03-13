# Start a new minimal image for running the Go application
FROM gcr.io/distroless/base-debian10

WORKDIR /root/

# Copy the binary from the builder image
COPY main .

# Expose the HTTP port
EXPOSE 8080

# Run the binary
CMD ["/root/main"]
