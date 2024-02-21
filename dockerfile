FROM golang:1.18-alpine3.16 AS builder
RUN mkdir /work
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./*.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /work/app /
CMD ["/app"]
