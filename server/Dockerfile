FROM golang:1.22 
# AS build-stage

WORKDIR /app
COPY . /app/.
RUN go mod download


# RUN GOOS=linux go build -o avito-banner-service main.go
#
# FROM gcr.io/distroless/base-debian11 AS build-release-stage
#
# WORKDIR /
#
# COPY --from=build-stage /app/avito-banner-service /avito-banner-service

EXPOSE 8080
CMD ["/app/avito-banner-service"]
