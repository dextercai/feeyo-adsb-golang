FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /workdir/

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN apk add git make
RUN make -f ci/makefile docker

FROM alpine:latest
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata wget curl
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/feeyo-adsb-golang /app/feeyo-adsb-golang

CMD ["/app/feeyo-adsb-golang", "-conf=/app/conf.ini"]