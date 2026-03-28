FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o compute .

FROM scratch

COPY --from=builder /app/compute /

ENTRYPOINT ["/compute"]
CMD ["0.0", "0.0"]
