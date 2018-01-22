#!/bin/bash

mkdir -p $GOPATH/src/github.com/tensorflow
cp -R serving/ $GOPATH/src/github.com/tensorflow/
cp -R tensorflow/ $GOPATH/src/github.com/tensorflow/