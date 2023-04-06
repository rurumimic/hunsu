# Hello

- doc: [quickstart](https://docs.lmql.ai/en/latest/quickstart.html)

## Setup

### Set Python

- rurumimic: [python](https://github.com/rurumimic/supply/blob/master/languages/python.md)

```bash
pyenv virtualenv 3.11.1 lmql
pyenv activate lmql
pip install --upgrade pip setuptools wheel
pip install black
```

### Install LMQL

- doc: [install](https://docs.lmql.ai/en/latest/installation.html)

```bash
pip install lmql
```

### OpenAI API Key

OpenAI: [API Keys](https://platform.openai.com/account/api-keys)

1. Create a new secret key.
2. Save a key: `.env`

```bash
mkdir -p $HOME/.lmql
```

```bash
# vi $HOME/.lmql/api.env
openai-org: <org identifier>
openai-secret: <api secret>
```

---

## Code

### Hello World

- [hello.py](hello.py)
