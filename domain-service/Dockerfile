FROM golang:1.11.5-alpine3.9

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/port-project/domain-service

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN apk add --no-cache git

# Download all the dependencies
RUN go get -d -v "google.golang.org/grpc"
RUN go get -d -v "github.com/golang/protobuf/proto"
RUN go get -d -v "gopkg.in/natefinch/lumberjack.v2"

# Install the package
RUN go install -v "gopkg.in/natefinch/lumberjack.v2"
RUN go install -v "google.golang.org/grpc"
RUN go install -v "github.com/golang/protobuf/proto"

RUN go build -o domain-service .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["./domain-service"]