#!/bin/bash
set -e

echo "Generating pki-vault-service package"
protoc --go_out=pkivaultrpc --go_opt=paths=source_relative  --go-grpc_out=pkivaultrpc --go-grpc_opt=paths=source_relative pkivault.proto

echo "cleaning modules"
go mod tidy

set +e
echo "Building package"
goBuildResult=$(go build -a -installsuffix cgo -o app . 2>&1)
if [[ "$?" != "0" ]]; then
    echo -e "\e[31m!!Building package: fail \e[0m"
    echo -e "$goBuildResult"
    exit 
fi