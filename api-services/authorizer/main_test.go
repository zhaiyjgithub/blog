package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_calcApiKey(t *testing.T) {
	os.Setenv("salt", "~!@9&6_!@$@56d")
	var appName = "cool-box"
	key, _ := calcApiKey(appName)
	fmt.Println(key)
}
