#!/bin/bash

# generates go file using proto file
protoc sumpb/sum.proto --go_out=plugins=grpc:.