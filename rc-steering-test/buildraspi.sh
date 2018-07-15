#!/bin/bash
cd ~/go/src/github.com/RyAndrew/raspi-embd/rc-steering-test/

export GOARCH=arm
go build

scp -i ~/pikey rc-steering-test pi@10.88.0.115:~/rc-steering-test/

scp -r -i ~/pikey webroot pi@10.88.0.115:~/rc-steering-test/