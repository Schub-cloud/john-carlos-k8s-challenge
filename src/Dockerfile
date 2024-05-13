FROM golang:1.22.3 as builder

WORKDIR /app
COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM alpine:latest
RUN apk --no-cache add ca-certificates


EXPOSE 80

COPY --from=builder /app/app /app/app

CMD ["/app/app"]