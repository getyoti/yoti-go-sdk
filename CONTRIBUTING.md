# Contributing

The command `go get "github.com/getyoti/yoti-go-sdk/v2"` downloads the Yoti package, along with its dependencies, and installs it.

## Commit Process

This repo comes with pre-commit hooks. We strongly recommend installing them with `pre-commit install`. This will lint and run unit tests automatically

## VS Code

For developing in VS Code, use the following `launch.json` file (placed inside a `.vscode` folder) to easily run the examples from VS Code:

```javascript
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
    {
        "name": "AML Example",
        "type": "go",
        "request": "launch",
        "mode": "debug",
        "program": "${workspaceFolder}/examples/aml/main.go"
    },
    {
        "name": "Example",
        "type": "go",
        "request": "launch",
        "mode": "debug",
        "remotePath": "",
        "host": "127.0.0.1",
        "program": "${workspaceFolder}/examples/profile/main.go",
        "env": {},
        "args": ["certificatehelper.go"],
        "showLog": true
    }
    ]
}
```
