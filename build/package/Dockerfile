FROM golang:1.22-alpine3.19 as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./cmd ./cmd
COPY ./internal ./internal

WORKDIR /app/cmd
RUN go build -o server


FROM alpine:3.15

COPY --from=build /app/cmd/server ./server

EXPOSE 8180

ENTRYPOINT ["./server"]


