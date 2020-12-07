#!/usr/bin/env bash

mkdir dto

protoc -I ecust_library_idl --gofast_out=dto ecust_library_idl/*.proto
