# build
FROM golang:1.20 AS build

WORKDIR /app
COPY * /app

RUN go mod download \
    && go build 

# runtime
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /app/go-echo /go-echo

EXPOSE 3000

ENTRYPOINT ["/go-echo"]
