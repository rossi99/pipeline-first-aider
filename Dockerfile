FROM golang:1.24.4
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./
RUN go mod download

# Copy necessary code into the container
COPY . ./

# Build the first-aider binary
RUN go build -o first-aider ./cmd

ENTRYPOINT ["/app/first-aider"]