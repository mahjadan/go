FROM golang:1.17.2-alpine3.14 as builder

RUN mkdir /go-proto-poc
WORKDIR /go-proto-poc

RUN apk add --update --no-cache git bash gcc musl-dev openssh-client util-linux curl

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  -o go-poc go-proto/main.go

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go-proto-poc/go-poc .

CMD ["./go-poc"]

#RUN go get -d github.com/cosmtrek/air && go get -d gotest.tools/gotestsum

EXPOSE 8888
