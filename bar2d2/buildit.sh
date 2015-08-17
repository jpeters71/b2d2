#!/bin/bash

pushd dto
go install 
popd

pushd persistence
go install
popd

pushd services
go install 
popd

pushd ui
go build
go install 
popd


echo "Build done."

