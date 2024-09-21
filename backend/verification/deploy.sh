#!/bin/bash
docker build -t mcaxdev/verification . && \
docker-compose up -d
