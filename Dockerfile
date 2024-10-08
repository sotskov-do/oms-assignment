FROM golang:1.22.5-alpine3.19 as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o application ./cmd/main.go

FROM alpine:3.19

COPY --from=builder /app/application /app/application
CMD ["/app/application"]