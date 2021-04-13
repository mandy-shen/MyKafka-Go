FROM scratch AS scratch
WORKDIR /go
COPY go.mod .
COPY go.sum .
COPY kafka.go .

FROM golang:alpine AS stage1
WORKDIR /build
COPY --from=scratch /go .
# smaller binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' kafka.go

FROM openjdk:8-jre-alpine
COPY kfk_start.sh .
RUN apk add --no-cache bash tar \
    && chmod +x kfk_start.sh \
    && cp kfk_start.sh /usr/bin \
    && rm kfk_start.sh
ADD kafka_2.12-2.7.0.tgz /
RUN apk del tar
COPY --from=stage1 /build/kafka /kafka
ENTRYPOINT [ "/kafka" ]
CMD [ "add-row neu-student, My Name, ID0000001" ]