FROM golang:alpine as builder

RUN apk add --no-cache git
ADD . /go/src/github.com/plutov/surv/survey
WORKDIR /go/src/github.com/plutov/surv/survey

# Install dependencies
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main github.com/plutov/surv/survey/cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/plutov/surv/survey/main .
COPY --from=builder /go/src/github.com/plutov/surv/survey/pkg/api/swagger.json /root/swagger.json

ENTRYPOINT [ "./main" ]

EXPOSE 8080