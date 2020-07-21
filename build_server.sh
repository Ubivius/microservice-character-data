#!/bin/bash

echo "Validating"
swagger validate swagger.yml
echo ""
echo "Generating server"
swagger generate server -A characterdata -f ./swagger.yml
echo ""
echo "Building and installing server"
go install ./cmd/characterdata-server/
echo ""
echo "Serving UI"
# swagger serve swagger.yml
