#!/bin/bash
export GOARCH=arm
go build pca9685-servo-webtest.go

scp -i ~/pikey pca9685-servo-webtest pi@10.88.0.115:~/
