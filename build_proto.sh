#!/usr/bin/env bash

mkdir dto

protoc -I proto --gofast_out=dto proto/*.proto
