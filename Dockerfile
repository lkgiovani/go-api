FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download && go mod tidy

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-api ./cmd/start

# Fase final
FROM alpine:3.20.3

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/go-api .
# Copia o diretório de migrações para a imagem final
COPY --from=builder /app/internal/app/infra/config/db/migrations /app/internal/app/infra/config/db/migrations

EXPOSE 8181

ENTRYPOINT ["./go-api"]