FROM golang:1.21-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api main.go

## scratch - menor imagem docker
FROM scratch
COPY --from=builder /app/api /
EXPOSE 8080:8080
CMD [ "/api" ]