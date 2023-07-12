FROM golang:alpine as build
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Fetch dependencies.
RUN go get -d -v
# Build the binary.
RUN go build -o /opt/moviedb/moviedb


FROM scratch
COPY --from=build /opt/moviedb/moviedb /opt/moviedb/moviedb
ENV MONGO_URI "mongodb://localhost:27017"
ENV SERVER_HOST_PORT "0.0.0.0:8090"
EXPOSE 8090
ENTRYPOINT ["/opt/moviedb/moviedb"]

