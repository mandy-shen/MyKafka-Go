FROM golang:alpine AS stage1
WORKDIR /build
COPY . .
# smaller binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' kafka.go

FROM openjdk:8-jre-alpine
WORKDIR /
COPY kfk_start.sh .
RUN apk add --no-cache bash tar \
    && chmod +x kfk_start.sh \
    && cp kfk_start.sh /usr/bin
ADD kafka_2.12-2.7.0.tgz /
COPY --from=stage1 /build/kafka /kafka
RUN apk del tar
ENTRYPOINT [ "/kafka" ]
CMD [ "add-row neu-student, My Name, ID0000001" ]