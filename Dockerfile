FROM golang:1.15.6-alpine3.12

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/brkelkar/smart_pharmacy_api

RUN apk --no-cache add git 

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .


# Install the package
RUN go build -tags=jsoniter .
# This container exposes port 8080 to the outside world
EXPOSE 8083

# Run the executable
CMD ["./smart_pharmacy_api"]