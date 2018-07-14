#!/bin/bash
cd ~/go/src/github.com/RyAndrew/raspi-embd/pca96850-servo-webtest/

export GOARCH=arm
go build pca9685-servo-webtest.go

scp -i ~/pikey pca9685-servo-webtest pi@10.88.0.115:~/pca9685-servo-webtest/

#scp -r -i ~/pikey webroot pi@10.88.0.115:~/pca9685-servo-webtest/