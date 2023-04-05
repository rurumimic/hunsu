# Speech

- language
  - go: [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai)

## API Key

OpenAI: [API Keys](https://platform.openai.com/account/api-keys)

1. Create a new secret key.
2. Save a key: `.env`

---

## Dev Container

- vscode: [containers](https://code.visualstudio.com/docs/devcontainers/containers)
  - feature: [apt-get-packages](https://github.com/devcontainers-contrib/features/tree/main/src/apt-get-packages)

### Setup .devcontainer/devcontainer.json

1. `F1`: Dev Containers: Open Folder in Container...
2. Go: `1.20`
3. Select: `apt-get packages (for Debian/Ubuntu)`
4. Keep Default

### Add a package

```json
{
  "name": "Go",
  "image": "mcr.microsoft.com/devcontainers/go:0-1.20",
  "features": {
    "ghcr.io/devcontainers-contrib/features/apt-get-packages:1": {
      "packages": ["portaudio19-dev"]
    }
  }
}
```

### Start a container

1. `F1`: Dev Containers: Open Folder in Container...
1. `Control` + `Shift` + `

```bash
pwd # /workspaces/hunsu/src/openai/go/speech
whoami # vscode
go version # go version go1.20.2 linux/amd64
```

---

## Start

### Init a mod

```bash
go mod init speech
```

### Install a library

```bash
go get github.com/joho/godotenv
go get -u github.com/MarkKremer/microphone
go get github.com/sashabaranov/go-openai
```

### main.go

- [main.go](main.go)

### Hello, AI

```bash
go run . record.wav
```

```bash
```
