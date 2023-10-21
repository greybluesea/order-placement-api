# Build stage
FROM golang:1.21.3 AS builder

WORKDIR /usr/src/app

# Copy only the Go module files to the build stage for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application source code
COPY . .




# Change to the /cmd directory
WORKDIR /usr/src/app/cmd

# Build your Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o order-placement-api .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage to the final image
COPY ./great .
COPY --from=builder /usr/src/app/cmd/order-placement-api .

# You may need to add additional dependencies or configurations here
EXPOSE 3000

# Start your application
CMD ["./order-placement-api"]
