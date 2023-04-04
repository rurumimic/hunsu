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
Welcome to the OpenAI Chatbot!
Type 'exit' to quit.

You: 안녕?
Bot: 안녕하세요! 저는 인공지능 어시스턴트입니다. 무엇을 도와드릴까요? :)

You: 100만이랑 10Billion 중에 더 큰 숫자가 뭐야?
Bot: 10 Billion이 더 큰 숫자입니다. 

100만은 1,000,000이지만, 
10 Billion은 10,000,000,000입니다. 
즉, 10 Billion은 100만보다 10,000배나 큰 숫자입니다.

You: exit
```
