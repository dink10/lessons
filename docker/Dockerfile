FROM golang:1.16

RUN go env

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/dink10/app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Tests
RUN go test ./...

# Download all the dependencies
RUN go mod tidy

# Download all the dependencies
RUN go build main.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]