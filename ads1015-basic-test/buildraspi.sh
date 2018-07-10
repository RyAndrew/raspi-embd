#!/bin/bash
export GOARCH=arm
go build ads1015-basic-test.go

scp -i ~/pikey ads1015-basic-test pi@10.88.0.115:~/
