package main

import (
	"os"
	"testing"
)

func Test_sendEmail(t *testing.T) {
//email_sender: chuck.zhaiyj@gmail.com
//email_sender_app_password: isnejfxzaniutstq
	os.Setenv("email_sender", "chuck.zhaiyj@gmail.com")
	os.Setenv("email_sender_app_password", "isnejfxzaniutstq")
	sendEmail([]string{"yuanji.zhai@mdland.com"}, "Hello", "Hi, test message.")
}
