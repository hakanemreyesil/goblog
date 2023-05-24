# FROM golang:alpine3.17
# WORKDIR /go/src/goweb
# COPY . .
# CMD ["/go/src/goweb/goblog"]

# First stage: build the application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Second stage: use the built artifact
FROM alpine:latest  

# Set the working directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port for the app (replace with your port)
EXPOSE 8080

# Command to run the binary
CMD ["./main"] 