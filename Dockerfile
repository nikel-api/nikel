# Should be latest stable release
FROM golang:1.15

# Setup working directory
WORKDIR /go/src/app
COPY . .

# Get & install dependencies, then build
RUN go get -d -v ./...
RUN go install -v ./...

# Set to whatever we're listening on
EXPOSE 8080

ENTRYPOINT ["nikel-core"]