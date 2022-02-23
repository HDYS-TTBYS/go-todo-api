FROM golang:alpine as build-env

RUN apk add --no-cache git gcc musl-dev
ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo
WORKDIR /go/src/github.com/HDYS-TTBYS/go-todo-api
COPY . .
RUN go build main.go

FROM alpine:3.14
WORKDIR /go/bin
ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo
COPY --from=build-env /go/src/github.com/HDYS-TTBYS/go-todo-api/main /
CMD [ "/main" ]
