# Build-Stage
FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go generate ./... &&  go build -o api ./cmd/api

FROM gcr.io/distroless/base
COPY --from=build /app/assets /assets
COPY --from=build /app/api /api
CMD ["/api"]
