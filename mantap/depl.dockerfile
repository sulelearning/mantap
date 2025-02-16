FROM golang:1.23-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/Zulhaidir/microservice
COPY go.mod go.sum ./
# RUN go mod download
COPY mantap mantap
RUN GO111MODULE=on go build -o /go/bin/app ./mantap/cmd/mantap \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["sleep", "infinity"]