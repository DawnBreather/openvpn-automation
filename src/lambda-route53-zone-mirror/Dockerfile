FROM golang as build
COPY . ./

RUN go get github.com/aws/aws-sdk-go/aws

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o route53-zone-mirror main/main.go

FROM ubuntu:18.04
RUN apt update \
    && apt install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build /go/route53-zone-mirror /bin
ENTRYPOINT route53-zone-mirror
