# build stage
FROM golang:1.23-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Zulhaidir/microservice
COPY . .
RUN go build -o /go/bin/app ./mantap/cmd/mantap \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# run stage
FROM alpine:3.13
WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY mantap/start.sh .
COPY mantap/app.env .
COPY mantap/db/migration ./migration
RUN chmod +x start.sh
EXPOSE 8080
CMD ["app"]
ENTRYPOINT ["./start.sh"]