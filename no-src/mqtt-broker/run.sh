#!/bin/bash

docker run -d \
    -p 1883:1883 \
    -p 9001:9001 \
    --name mqtt-broker \
    -v /mosquitto/data \
    -v /mosquitto/log \
    -v /home/nkuhn/go/src/iteragit.iteratec.de/traze/goclient/no-src/mqtt-broker/mosquitto.conf:/mosquitto/config/mosquitto.conf \
    eclipse-mosquitto

# docker run -d -p 1883:1883 -p 9001:9001 --name mqtt-broker -v /mosquitto/data -v /mosquitto/log eclipse-mosquitto
# docker cp mosquitto.conf mqtt-broker:/mosquitto/config/