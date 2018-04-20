FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/brunobandev/tournament-api/api

# Install our dependencies
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/gorilla/mux

# Install api binary globally within container 
RUN go install github.com/brunobandev/tournament-api/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (4001)
EXPOSE 4001 