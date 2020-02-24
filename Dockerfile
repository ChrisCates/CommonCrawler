# Start from golang v1.11 base image
FROM golang:alpine AS build

ENV GO111MODULE=on

# Maintainer info
LABEL maintainer="Chris Cates <hello@chriscates.ca>, Onuwa Nnachi Isaac <matrix4u2002@gmail.com>"

# Set current working directory inside the container
WORKDIR /app

# Copy everything from the source directory to destination directory inside the container
COPY . .

# Download all dependencies
RUN go mod download

# build the package
RUN go build -o common-crawler main.go

## build the smaller image
FROM alpine

LABEL maintainer="Chris Cates <hello@chriscates.ca>, Onuwa Nnachi Isaac <matrix4u2002@gmail.com>"

WORKDIR /app

COPY --from=build /app/common-crawler .
COPY wet.paths .

RUN apk add ca-certificates && \
    chmod +x common-crawler

CMD ["./common-crawler"]
