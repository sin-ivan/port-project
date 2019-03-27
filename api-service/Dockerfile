
# Start from golang alpine
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/sin-ivan/port-project/api-service/

# Copy everything from the current directory to the working directory inside the container
COPY . .

RUN apk add --no-cache git

# Download all the dependencies
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/api-service .

######## Start a new stage from scratch #######
FROM scratch  

EXPOSE 8080

COPY --from=builder /go/bin/api-service .
COPY --from=builder /go/src/github.com/sin-ivan/port-project/api-service/resources/ports.json /resources/

CMD ["./api-service"]