FROM golang:1.15 AS build
WORKDIR /go/src

COPY pkg ./pkg
COPY go.mod .
COPY go.sum .
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o pager-cli-client .

FROM scratch AS runtime

COPY --from=build /go/src/pager-cli-client ./
ENTRYPOINT ["./pager-cli-client"]

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
