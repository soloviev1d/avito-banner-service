FROM golang:1.22-bookworm

COPY ./ ./
RUN go mod download


RUN GOOS=linux go build -o bin/avito-banner-service cmd/main.go


EXPOSE 8080
CMD ["./bin/avito-banner-service"]
