FROM golang:latest
COPY . /go/src/gener
# ENV GOPATH /go
# ENV PATH $GOPATH/bin:$PATH
WORKDIR /go/src/gener

#RUN go get -u github.com/golang/dep/...
#RUN dep ensure

# RUN go run main.go
RUN go build -o /app/gener .
CMD ["/app/gener"]
EXPOSE 8085
