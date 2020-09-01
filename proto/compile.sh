#!/usr/bin/env sh

# Install proto3 from source
#  brew install autoconf automake libtool
#  git clone https://github.com/google/protobuf
#  ./autogen.sh ; ./configure ; make ; make install
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples
# Or install protoc and go plugin with Debian-based distros and APT
# sudo apt install libprotobuf-dev protobuf-compiler golang-goprotobuf-dev -y

# Golang compile
# Requires GOROOT, GOPATH and  GOBIN env variables
protoc -I . ./identity.proto --go_out=plugins=grpc:.