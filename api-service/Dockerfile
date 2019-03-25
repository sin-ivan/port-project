FROM golang:1.11.5-alpine3.9

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/port-project/api-service

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

RUN go build -o api-service .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["./api-service"]


#---------------------
# FROM golang:alpine
# ADD ./src /go/src/app
# WORKDIR /go/src/app

# RUN apk add git

# RUN go get -d -v "github.com/gorilla/mux"
# RUN go get -d -v "gopkg.in/natefinch/lumberjack.v2"

# RUN go install -v "github.com/gorilla/mux"
# RUN go install -v "gopkg.in/natefinch/lumberjack.v2"

# ENV PORT=8080
# EXPOSE 8080

# CMD ["go", "run", "main.go"]
#-------------------------------

# FROM golang:alpine AS builder
# RUN go version

# WORKDIR /app

# ENV SRC_DIR=/go/src/github.com/port-project/api-service/

# ADD . $SRC_DIR

# RUN apk add --no-cache git

# RUN go get -d -v "gopkg.in/natefinch/lumberjack.v2"
# RUN go install -v "gopkg.in/natefinch/lumberjack.v2"

# RUN go get -d -v "google.golang.org/grpc"
# RUN go install -v "google.golang.org/grpc"

# RUN go get -d -v "github.com/golang/protobuf/proto"
# RUN go install -v "github.com/golang/protobuf/proto"

# # Build it:
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd $SRC_DIR; go build -o myapp; cp myapp /app/

# # ENV PORT=8080
# EXPOSE 8080

# ENTRYPOINT ["./myapp"]

# ADD ./src /go/src/app
# WORKDIR /go/src/app

# RUN apk add --no-cache git

# RUN go get -d -v "gopkg.in/natefinch/lumberjack.v2"
# RUN go install -v "gopkg.in/natefinch/lumberjack.v2"

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#



# FROM scratch
# WORKDIR /root/
# COPY --from=builder /go/src/app .

# ENV PORT=8080
# EXPOSE 8080
# ENTRYPOINT ["./main"]
