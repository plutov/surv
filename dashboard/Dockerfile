FROM golang:alpine as builder

RUN apk add --no-cache git
ADD . /go/src/github.com/plutov/surv/dashboard
WORKDIR /go/src/github.com/plutov/surv/dashboard

# Install dependencies
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main github.com/plutov/surv/dashboard/cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/plutov/surv/dashboard/main .
COPY --from=builder /go/src/github.com/plutov/surv/dashboard/pkg/api/swagger.json /root/swagger.json
COPY --from=builder /go/src/github.com/plutov/surv/dashboard/config.json /root/config.json

ENTRYPOINT [ "./main" ]

EXPOSE 8080