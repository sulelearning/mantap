FROM golang:1.23-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Zulhaidir/microservice
COPY go.mod go.sum ./
COPY mantap mantap
RUN go mod download \
    && GO111MODULE=on go build -o /go/bin/app ./mantap/cmd/mantap \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]