FROM golang:alpine

RUN apk add build-base

ENV PORT=8080

WORKDIR /go/src/app
COPY . .

# Go dependencies
# Migrate tool (CLI wrapper)
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/

# The rest of dependencies
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE $PORT

ENTRYPOINT [ "/bin/sh", "./start.sh" ]
