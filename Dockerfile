FROM golang:alpine

# ENV GIN_MODE=release
ENV PORT=8080

# TODO Copy only what's necessary
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE $PORT

ENTRYPOINT [ "go", "run", "app.go" ]
