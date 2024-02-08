# Go Lunch [![Quality assurance](https://github.com/mishamyrt/go-lunch/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/mishamyrt/go-lunch/actions/workflows/quality-assurance.yaml) [![GitHub release](https://img.shields.io/github/v/tag/mishamyrt/go-lunch)](https://GitHub.com/mishamyrt/go-lunch/releases/)

<img src="./docs/logo@2x.png" align="right" width="128" />

A library for easily creating launch agent files to autorun your applications in macOS.

- **Easy interface**
- **Pure Golang**
- **Safe**

## Installation

Run command at your project directory.

```sh
go get -u github.com/mishamyrt/go-lunch
```

## Usage

Create `Agent` instance with required fields: `PackageName`, `Command`, `Path`. For path you can use `UserPath` and `SharedPath` helpers. It returns launch agent path.

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

### Parameters

- `PackageName` - App package name
- `KeepAlive` - Restart app on exit
- `Command` - App start command
- `StandardOutPath` - App stdout. Defaults to /dev/null
- `StandardErrorPath`- App stderr. Defaults to /dev/null
- `Path` - Launch Agent path
