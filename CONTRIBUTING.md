# Contributing

The command `go get "github.com/getyoti/yoti-go-sdk/v2"` downloads the Yoti package, along with its dependencies, and installs it.

## Commit Process

1) `go build` builds the package (then discards the results)
1) `goimports` formats the code and sanitises imports
1) `go vet` reports suspicious constructs
1) `go test -race` to run the tests and detect race conditions
1) `golangci-lint run` for [GolangCI-Lint](https://github.com/golangci/golangci-lint)
1) `go mod tidy` prunes any no-longer-needed dependencies from `go.mod`, and adds any dependencies needed

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