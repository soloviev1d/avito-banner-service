FROM golang:1.22

WORKDIR /app


COPY ./ ./

RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /avito-banner-service


EXPOSE 8080

CMD ["/avito-banner-service"]
