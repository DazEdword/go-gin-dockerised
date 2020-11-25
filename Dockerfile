FROM golang:alpine

RUN apk add build-base

# ENV GIN_MODE=release
ENV PORT=8080

# TODO Copy only what's necessary
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE $PORT

CMD [ "go", "run", "app.go" ]
