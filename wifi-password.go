package wifipassword

import (
	"os"
	"runtime"
	"os/exec"
	"github.com/yelinaung/wifi-name"
)

func WifiPassword() string {
	var wifiName = wifiname.WifiName()

	// I am aware of this, but I have no idea for now
	// Since, I am still learing. Will do soemthing later
	if wifiName == "Could not get SSID" {
		return "You still haven't connected to your wifi!"
	}

	platform := runtime.GOOS

	if platform == "darwin" {
		return forOSX(wifiName)
	} else if platform == "win32" {
		// TODO for Windows
		return ""
	} else {
		// TODO for Linux
		return ""
	}
}

func forOSX(wifiName string) string{
	var (
		cmdOut []byte
		err    error
	)

	cmdName := "security"
	cmdArgs := []string{"find-generic-password", "-wga", wifiName}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		panic(err)
		os.Exit(1);
	}

	return string(cmdOut)
}