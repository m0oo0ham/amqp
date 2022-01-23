#!/bin/bash

docker run -d --rm --name mq -p 8161:8161 -p 5672:5672 rmohr/activemq