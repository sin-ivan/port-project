# FROM golang:1.11.5-alpine3.9

# # Set the Current Working Directory inside the container
# WORKDIR $GOPATH/src/github.com/port-api-project/go-docker

# # Copy everything from the current directory to the PWD(Present Working Directory) inside the container
# COPY . .

# RUN apk add --no-cache git

# # Download all the dependencies
# # https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
# RUN go get -d -v "github.com/gorilla/mux"
# RUN go get -d -v "gopkg.in/natefinch/lumberjack.v2"

# # Install the package
# RUN go install -v "github.com/gorilla/mux"
# RUN go install -v "gopkg.in/natefinch/lumberjack.v2"

## RUN go build -o go-docker .

# # This container exposes port 8080 to the outside world
# EXPOSE 8080

# # Run the executable
# CMD ["main"]


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

FROM golang:1.11.5 AS builder
RUN go version

ADD ./src /go/src/app
WORKDIR /go/src/app

RUN go get -d -v "github.com/gorilla/mux"
RUN go get -d -v "gopkg.in/natefinch/lumberjack.v2"

RUN go install -v "github.com/gorilla/mux"
RUN go install -v "gopkg.in/natefinch/lumberjack.v2"
# RUN set -x && \
#     go get github.com/golang/dep/cmd/dep && \
#     dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/app .

EXPOSE 8080
ENTRYPOINT ["./main"]
