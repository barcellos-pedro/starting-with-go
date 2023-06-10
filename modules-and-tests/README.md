# Modules

- Create a folder and a file for the module

```go
 package foo

 func Say() string{
    return "Hey"
 }
```

- Run **go mod init [module-name]**

```bash
    go mod init modules/foo
```

- In another folder to use the module, run **go mod init [module-name]**

```bash
    go mod init modules
```

- Import the modules in the main file

```go
    package main

    import (
	    "fmt"
	    "modules/foo"
    )

    func main() {
        fmt.Println(foo.Say())
    }
```

- Run go mod tidy

- If necessary for local modules run

```bash
go mod edit -replace modules/foo=./foo
```
