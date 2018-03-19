# Contributing

The command `go get "github.com/getyoti/yoti-go-sdk"` downloads the Yoti package, along with its dependencies, and installs it.

## Commit Process

1) `go fmt` to format the code
1) `go vet` reports suspicious constructs
1) `go test` to run the tests

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
        "args": [],
        "showLog": true
    }
    ]
}
```