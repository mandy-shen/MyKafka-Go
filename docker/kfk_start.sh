#!/bin/bash
export KAFKA_HOME=/kafka_$KAFKA_VER
$KAFKA_HOME/bin/zookeeper-server-start.sh -daemon $KAFKA_HOME/config/zookeeper.properties \
&& until echo ruok | nc 127.0.0.1 2181 > /dev/null; do sleep 1; done \
&& $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties