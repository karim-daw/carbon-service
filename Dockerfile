# Build stage
FROM golang:1.22.1 as builder

# Set the working directory
WORKDIR /go/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# Final stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/app/main .

# Use a non-root user
RUN adduser -D user
USER user

# Command to run the executable
CMD ["./main"]
