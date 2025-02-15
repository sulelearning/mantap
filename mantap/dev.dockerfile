FROM golang:1.23-alpine AS build

RUN apk --no-cache add gcc g++ make ca-certificates

# Set working directory to the root
WORKDIR /mnt/microservice/mantap

# Menyalin go.mod dan go.sum ke dalam kontainer
COPY ../go.* ./

# Mengunduh dependensi
RUN go mod download

# Menyalin seluruh kode aplikasi ke dalam kontainer
COPY ../ ./

# Menginstal air dan goimports untuk hot-reloading dan format
RUN go install github.com/air-verse/air@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest


# Mengekspos port yang digunakan aplikasi
EXPOSE 8080

# Menjalankan aplikasi dengan air
CMD ["air", "-c", ".air.toml"]