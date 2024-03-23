# Use base Ubuntu image
FROM ubuntu:latest
RUN apt-get update && apt-get install -y netcat
# Copy the Go binary into the container
COPY ./leaf-analyzer-server /app/
# Make the binary executable
RUN chmod +x /app/leaf-analyzer-server
CMD ["/app/leaf-analyzer-server"]
