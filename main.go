package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func getExternalIP() (string, error) {
	res, err := http.Get("http://checkip.amazonaws.com/")
	if err != nil {
		return "", err
	}
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", string(ip)), nil

}

func getInternalIP() (string, error) {
	cmd := exec.Command("ipconfig", "getifaddr", "en0")
	res, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", string(res)), nil
}

func main() {
	eip, err := getExternalIP()
	if err != nil {
		panic(err)
	}
	fmt.Printf("IP PUBLIC: %v\n", eip)
	iip, err := getInternalIP()
	if err != nil {
		panic(err)
	}
	fmt.Printf("IP INTERNAL: %v\n", iip)

}
