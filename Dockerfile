FROM golang:1.15.6-alpine3.12

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/kaleganeshrajan/smart_pharmacy_api

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .


ENV  GIN_MODE=release
RUN go build -tags=jsoniter .

# Run the executable
CMD ["./smart_pharmacy_api"]