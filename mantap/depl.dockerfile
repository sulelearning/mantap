FROM golang:1.23-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Zulhaidir/microservice
COPY go.mod go.sum ./
COPY mantap mantap
RUN GO111MODULE=on go build -o /go/bin/app ./mantap/cmd/mantap \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]