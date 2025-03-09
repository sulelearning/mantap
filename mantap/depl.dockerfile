# build stage
FROM docker.io/library/golang:1.23.3-alpine3.19 AS build
WORKDIR /go/src/microservice
COPY go.mod go.sum ./
COPY mantap mantap
RUN go build -o /go/bin/app ./mantap/cmd/mantap && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# run stage
FROM docker.io/library/alpine:3.19
WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY mantap/start.sh .
COPY mantap/app.env .
COPY mantap/db/migration ./migration
RUN chmod +x start.sh
EXPOSE 8080
CMD ["app"]
ENTRYPOINT ["./start.sh"]