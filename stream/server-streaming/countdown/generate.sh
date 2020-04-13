#!/bin/bash

# generates go file using proto file
protoc countdownpb/countdown.proto --go_out=plugins=grpc:.