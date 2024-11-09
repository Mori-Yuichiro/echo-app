FROM golang:1.23 AS base
WORKDIR /base
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /base/main.go

FROM golang:1.23 AS dev
WORKDIR /app
COPY --from=base /base .
ENV GO_ENV=dev
RUN go install github.com/air-verse/air@latest
CMD ["air"]

FROM alpine:latest AS production
WORKDIR /app
COPY --from=base /base/main .
COPY .env .
EXPOSE 8080
CMD ["/app/main"]