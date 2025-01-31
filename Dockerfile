#2
FROM golang:latest AS test-stage
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
CMD ["go", "test", "-v", "./..."]