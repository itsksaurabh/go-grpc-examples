#!/bin/bash

# generates go file using proto file
protoc greetpb/greet.proto --go_out=plugins=grpc:.