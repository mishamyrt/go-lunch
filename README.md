# Go Lunch [![Quality assurance](https://github.com/mishamyrt/go-lunch/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/mishamyrt/go-lunch/actions/workflows/quality-assurance.yaml) [![GitHub release](https://img.shields.io/github/v/tag/mishamyrt/go-lunch)](https://GitHub.com/mishamyrt/go-lunch/releases/)

<img src="./docs/logo@2x.png" align="right" width="128" />

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