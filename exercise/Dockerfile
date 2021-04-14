FROM ubuntu

RUN set -eux; \
    apt-get update; \
    apt-get install -y gosu; \
    rm -rf /var/lib/apt/lists/*; \
# verify that the binary works
    gosu nobody true

COPY entry.sh .
RUN chmod +x entry.sh
COPY entry.sh /usr/bin

ENTRYPOINT ["entry.sh"]