FROM golang:1.24.4
WORKDIR /app

# Copy source and mappings
COPY main.go .
COPY errors.json .
COPY go.mod .
COPY go.sum .

# Build the binary
RUN go build -o first-aider main.go

ENTRYPOINT ["/app/first-aider"]