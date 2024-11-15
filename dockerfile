FROM golang:alpine

COPY . /app
WORKDIR /app
# Untested docker image
CMD [ "go", "run", "cmd/api/main.go"]