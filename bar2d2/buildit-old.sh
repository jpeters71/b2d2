#!/bin/bash

go install ./dto
go install ./persistence
go install ./services

rm -f ./ui/ui
go build ./ui
go install ./ui 

echo "Build done."
