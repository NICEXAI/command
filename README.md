# command
exec-based command tools.

### Installation

Run the following command under your project:

> go get -u github.com/NICEXAI/command

### Basic Usage

If I want to execute the following command：

> go build -o ./example/example.exe ./example/example.go

you can:

```go
    if cmd, err := command.Run("go", "build", "-o", "./example/example.exe", "./example/example.go"); err != nil {
        log.Printf("command exec failed, error: %v", err)
        return
    }

    cmd.Wait()
    log.Printf("command exec success, error: %v", err)
```

If you want to stop the currently executing command, you can use `cmd.Stop()`, for example:

```go
    cmd, err = command.Run("./example/example.exe")
    if err != nil {
        log.Printf("command exec failed, error: %v", err)
        return
    }

    go func() {
        time.Sleep(10 * time.Second)
        log.Println("process begins to exit")
        cmd.Stop()
    }()

    cmd.Wait()
    log.Printf("process exit success, error: %v", err)
```