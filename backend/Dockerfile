# GO Builder
FROM golang:1.22 AS go-builder
WORKDIR /app
COPY backend/ .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# GO App Server (backend)
FROM alpine:3.14 AS go-server
WORKDIR /app
COPY --from=go-builder /app/server .
EXPOSE 8080
CMD ["./server"]
