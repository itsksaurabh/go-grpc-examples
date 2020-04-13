#!/bin/bash

# generates go file using proto file
protoc sumAllpb/sumAll.proto --go_out=plugins=grpc:.