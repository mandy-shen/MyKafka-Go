FROM scratch AS scratch
WORKDIR /go
COPY /go .

FROM golang:alpine AS stage1
WORKDIR /build
COPY --from=scratch /go .
# smaller binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' main.go

FROM openjdk:8-jre-alpine
ENV KAFKA_VER=2.12-2.7.0
COPY kfk_start.sh /usr/bin
COPY --from=stage1 /build/main /main
RUN apk add --no-cache bash tar netcat-openbsd \
    && chmod +x /usr/bin/kfk_start.sh
ADD kafka_$KAFKA_VER.tgz .
RUN apk del tar
ENTRYPOINT [ "/main" ]
CMD [ "add-row neu-student, My Name, ID0000001" ]