FROM golang:1.15
RUN mkdir -p /go/src/challenge
WORKDIR /go/src/challenge
COPY challenge/ .
CMD ["go", "test", "./..."]
