#!/bin/bash
export GOARCH=arm
go build pca9685-servotest.go
scp -i ~/pikey pca9685-servotest pi@10.88.0.115:~/