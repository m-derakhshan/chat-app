# Use official Go image
FROM golang:1.23.5

# Set working dir inside container
WORKDIR /app

COPY . .


RUN go mod download

# Copy everything else


# Build the app
RUN go build -o app ./cmd/main.go

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["./app"]
