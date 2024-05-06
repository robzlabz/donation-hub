FROM golang:1.20-alpine3.17 as build
WORKDIR /go/src/github.com/isdzulqor/donation-hub

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./cmd ./cmd
COPY ./internal ./internal

WORKDIR /go/src/github.com/isdzulqor/donation-hub/cmd
RUN go build -o server

FROM alpine:3.15
RUN apk add ca-certificates tzdata

COPY --from=build /go/src/github.com/isdzulqor/donation-hub/cmd/server ./server
COPY --from=build /go/src/github.com/isdzulqor/donation-hub/cmd/web ./web

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.8.1 /lambda-adapter /opt/extensions/lambda-adapter

ENV PORT=8180

ENTRYPOINT [ "./server" ]
