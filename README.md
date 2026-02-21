# grpc-template

## requirements

- go 1.26.0 or higher
- bufbuild/buf installed
- protoc-gen-go installed
- protoc-gen-go-grpc installed

## optional

- air installed

## setup

### install go

install go using following methods:

- macos:

    ```bash
    brew install go
    ```

- linux using `apt`:

    ```bash
    sudo add-apt-repository ppa:longsleep/golang-backports
    sudo apt update
    sudo apt install golang-go
    ```

- linux using `snap`:

    ```bash
    sudo snap install --classic go
    ```

then set up the environment variables:

- macos:

    ```bash
    echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.zshrc
    source ~/.zshrc
    ```

- linux:

    ```bash
    echo 'export PATH="$PATH:/usr/local/go/bin"' >> ~/.bashrc
    source ~/.bashrc
    ```

### install bufbuild/buf

```bash
brew install bufbuild/buf/buf
```

### install protoc-gen-go and protoc-gen-go-grpc

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.1
```

### install air

```bash
go install github.com/air-verse/air@v1.64.5
```

## features

- example template for grpc server
- tracing with opentelemetry
- logging with slog
- database migration and connection
- validation with go-playground/validator
- graceful shutdown
