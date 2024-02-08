// Example of app that creates or removes plist file to CWD
package main

import (
	"fmt"

	"github.com/mishamyrt/go-lunch"
)

const packageName = "com.domain.my-app"

func main() {
	agent := lunch.Agent{
		PackageName: packageName,
		Command:     "/usr/local/bin/my-app run -d",
		Path:        "./" + packageName + ".plist",
	}
	exist, err := agent.Exists()
	if err != nil {
		panic(err)
	}
	if exist {
		fmt.Println("Removing file")
		agent.Remove()
	} else {
		fmt.Println("Creating file")
		agent.Write()
	}
}
