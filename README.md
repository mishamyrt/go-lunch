# Go Lunch

<img src="./docs/logo.svg" align="right" width="70" />

A library for easily creating launch agent files to autorun your applications in macOS.

- **Easy interface**
- **Pure Golang**
- **Safe**

## Usage

```go
package main

import "github.com/mishamyrt/go-lunch"

const packageName = "com.domain.my-app"

func main() {
	path, _ := lunch.UserPath(packageName)
	agent := lunch.Agent{
		PackageName: packageName,
		Command:     "/usr/local/bin/my-app run -d",
		Path:        path,
	}
	agent.Write()
}
```