#!/bin/bash

pushd bar2d2/dto
go install 
popd

pushd bar2d2/persistence
go install
popd

pushd bar2d2/services
go install 
popd

pushd bar2d2/ui
go build
go install 
popd


echo "Build done."

