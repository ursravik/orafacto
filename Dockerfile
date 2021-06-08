FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir ./addition
RUN mkdir ./addsub
RUN mkdir ./addsubdiv

# Copy the code into the container
COPY addition.go ./addition
COPY addsub.go ./addsub
COPY addsubdiv.go ./addsubdiv

# Build the application
RUN go build -o addition ./addition
RUN go build -o addsub ./addsub
RUN go build -o addsubdiv ./addsubdiv


# Export necessary port
EXPOSE 3000

# Command to run when starting the container
#CMD ["./addition/addition"]
