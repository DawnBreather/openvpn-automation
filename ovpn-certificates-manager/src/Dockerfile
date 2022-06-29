FROM golang:latest as build-stage
WORKDIR /go/src/
COPY . .

RUN go get github.com/domodwyer/mailyak \
    && go get github.com/thoas/go-funk \
    && go get github.com/go-resty/resty

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o auth0-ovpn-users-sync main.go

FROM alpine
WORKDIR /opt/auth0-ovpn-users-sync
COPY --from=build-stage /go/src/auth0-ovpn-users-sync /opt/auth0-ovpn-users-sync/run-app
COPY templates ./templates

CMD ["./run-app"]
