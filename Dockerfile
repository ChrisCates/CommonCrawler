# Start from golang v1.11 base image
FROM golang

ENV GO111MODULE=on

# Maintainer info
LABEL maintainer="Onuwa Nnachi Isaac <matrix4u2002@gmail.com>"

# Set current working directory inside the container
WORKDIR /app

# Copy everything from the source directory to destination directory inside the container
COPY . .

# Download all dependencies
RUN go get -d -v ./...

# Install the packages
RUN go install -v ./...

RUN go run ./src/


# Expose port 8080 for the container to connect to outside world
EXPOSE 8080

# RUN the executable
CMD [ "CommonCrawler" ]
