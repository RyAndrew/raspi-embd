#!/bin/bash
cd ~/go/src/github.com/RyAndrew/raspi-embd/rc-steering-test-pid-gamepad

export GOARCH=arm
go build

#scp -i ~/pikey rc-steering-test-pid-gamepad pi@10.88.0.115:~/rc-steering-test-pid-gamepad/
scp -i ~/pi3key rc-steering-test-pid-gamepad pi@10.88.0.109:~/rc-steering-test-pid-gamepad

#scp -r -i ~/pikey webroot pi@10.88.0.115:~/rc-steering-test-pid-gamepad/
#scp -r -i ~/pi3key webroot pi@10.88.0.109:~/rc-steering-test-pid-gamepad/