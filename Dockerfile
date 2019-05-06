# Start from golang v1.11 base image
FROM golang

ENV GO111MODULE=on

# Maintainer info
LABEL maintainer="Chris Cates <hello@chriscates.ca>, Onuwa Nnachi Isaac <matrix4u2002@gmail.com>"

# Set current working directory inside the container
WORKDIR /app

# Copy everything from the source directory to destination directory inside the container
COPY . .

# Download all dependencies
RUN go get -d -v ./...

# Install and build the package
RUN go build -i -o ./dist/commoncrawler ./src/*.go

# Run the binary
CMD ["./dist/commoncrawler"]
