FROM golang:alpine as builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN apk update && apk add --update --no-cache alpine-sdk bash ca-certificates git

WORKDIR $GOPATH/src/apartment

COPY . .

RUN go mod download
RUN go build -ldflags="-w -s" -o /go/bin/apartment .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/apartment /go/bin/apartment

CMD ["/go/bin/apartment"]
