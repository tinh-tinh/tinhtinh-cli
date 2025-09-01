# CLI for Tinh Tinh

<div align="center">
<img alt="GitHub Release" src="https://img.shields.io/github/v/release/tinh-tinh/tinhtinh-cli">
<img alt="GitHub License" src="https://img.shields.io/github/license/tinh-tinh/tinhtinh-cli">
<a href="https://codecov.io/gh/tinh-tinh/tinhtinh-cli" > 
 <img src="https://codecov.io/gh/tinh-tinh/tinhtinh-cli/graph/badge.svg?token=VK57E807N2"/> 
 </a>
<a href="https://pkg.go.dev/github.com/tinh-tinh/tinhtinh-cli"><img src="https://pkg.go.dev/badge/github.com/tinh-tinh/tinhtinh-cli.svg" alt="Go Reference"></a>
</div>

<div align="center">
    <img src="https://avatars.githubusercontent.com/u/178628733?s=400&u=2a8230486a43595a03a6f9f204e54a0046ce0cc4&v=4" width="200" alt="Tinh Tinh Logo">
</div>

## Overview

**tinhtinh-cli** is the official command-line tool for Tinh Tinh, designed to streamline project creation and automation.

## Install

```
go install github.com/tinh-tinh/tinhtinh-cli/v2@latest
```

## Features

- **Project Initialization:** Instantly scaffold a new Tinh Tinh project with best practices.
- **Automated Boilerplate:** Generates main.go, modules, controllers, services, middleware, and guards with ready-to-use code.
- **Go Module Setup:** Handles `go mod init` and dependency management automatically.
- **Customizable:** Supports custom package names and optional Git repo cloning.

## Usage

```bash
tinhtinh-cli init my-service
```

This command will:
- Create a new project directory
- Initialize a Go module
- Generate a main entry point and starter app structure
- Scaffold controller and service templates

## Example Generated App

```go
package main

import (
    "github.com/tinh-tinh/tinhtinh/v2/core"
    "my-service/app"
)

func main() {
    server := core.CreateFactory(app.NewModule)
    server.Listen(3000)
}
```

## Contributing

We welcome contributions! Please feel free to submit a Pull Request.

## Support

If you encounter any issues or need help, you can:
- Open an issue in the GitHub repository
- Check our documentation
- Join our community discussions
