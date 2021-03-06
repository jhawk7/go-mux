# inherit from the Go official Image
FROM golang:1.8

# set a workdir inside docker
WORKDIR /go/src/github.com/dev/go-mux

# copy . (all in the current directory) to . (WORKDIR)
COPY . .

# run a command - this will run when building the image
RUN go get -u github.com/gorilla/mux
RUN go build -o go-mux

# the port we wish to expose
EXPOSE 8888

# run a command when running the container
CMD ./go-mux
