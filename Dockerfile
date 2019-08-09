FROM golang:1.12

WORKDIR /go/src/app

COPY . .

RUN go build -o app

COPY --from=build /go/src/app/app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]