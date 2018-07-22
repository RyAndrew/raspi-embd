#!/bin/bash
cd ~/go/src/github.com/RyAndrew/raspi-embd/rc-steering-test-pid

export GOARCH=arm
go build

scp -i ~/pikey rc-steering-test-pid pi@10.88.0.115:~/rc-steering-test-pid/

#scp -r -i ~/pikey webroot pi@10.88.0.115:~/rc-steering-test-pid/