FROM --platform=linux/arm64/v8 golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc g++ make git bash perl python3 curl pkgconfig musl-dev

# Install librdkafka
RUN git clone https://github.com/edenhill/librdkafka.git \
    && cd librdkafka \
    && ./configure --prefix /usr \
    && make \
    && make install

# Set up Go environment
ENV CGO_ENABLED=1
ENV CGO_LDFLAGS="-L/usr/lib -lrdkafka"
ENV CGO_CFLAGS="-I/usr/include"

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -tags musl -o main ./cmd/server

# Final stage
FROM --platform=linux/arm64/v8 alpine:latest

# Copy the built executable
COPY --from=builder /app/main /app/main

# Copy librdkafka shared libraries
COPY --from=builder /usr/lib/librdkafka*.so* /usr/lib/

# Set the entrypoint
ENTRYPOINT ["/app/main"]
