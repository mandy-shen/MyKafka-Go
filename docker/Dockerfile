FROM scratch AS scratch
COPY kfk_start.sh .

FROM openjdk:8-jre-alpine
ENV KAFKA_VER=2.12-2.7.0
COPY --from=scratch /kfk_start.sh /usr/bin
RUN apk add --no-cache bash tar netcat-openbsd \
    && chmod +x /usr/bin/kfk_start.sh
ADD kafka_$KAFKA_VER.tgz /
RUN apk del tar
ENTRYPOINT [ "/usr/bin/kfk_start.sh" ]
EXPOSE 2181 9092 19092