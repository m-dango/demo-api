FROM golang:1.25-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

FROM scratch
COPY --from=builder /build/main /bin/main
ENTRYPOINT ["/bin/main"]
