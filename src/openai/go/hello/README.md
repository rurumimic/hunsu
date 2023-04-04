# Hello

- language
  - go: [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai)

## API Key

OpenAI: [API Keys](https://platform.openai.com/account/api-keys)

1. Create a new secret key.
2. Save a key: `.env`

## Start

### Init a mod

```bash
go mod init hello
```

### Install a library

```bash
go get github.com/joho/godotenv
go get github.com/sashabaranov/go-openai
```

### main.go

- [main.go](main.go)

### Hello, AI

```bash
go run .
```

```bash
I'm an AI language model, so I don't have feelings, but I'm functioning properly and ready to assist you. How may I help you today?
```
