FROM golang AS builder

# copy the source
ADD .   /go/src/github.com/knative/docs/docs/serving/samples/gitwebhook-go
WORKDIR /go/src/github.com/knative/docs/docs/serving/samples/gitwebhook-go

# install dependencies
RUN go get github.com/google/go-github/github
RUN go get golang.org/x/oauth2
RUN go get gopkg.in/go-playground/webhooks.v3
RUN go get gopkg.in/go-playground/webhooks.v3/github

# build the sample
RUN CGO_ENABLED=0 go build -o /go/bin/webhook-sample .

FROM golang:alpine

EXPOSE 8080
COPY --from=builder /go/bin/webhook-sample /app/webhook-sample

ENTRYPOINT ["/app/webhook-sample"]
