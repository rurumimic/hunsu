# Custom

## Setup Env

### Python Env

Python v3.11.1

```bash
pyenv virtualenv 3.11.1 myllm
pyenv activate myllm
```

```bash
pip install --upgrade pip setuptools wheel
pip install black
pip install jupyterlab
```

### Install Library

```bash
pip install -r requirements.txt
```

---

## LLM Model

### git lfs

- [Git Large File Storage](https://git-lfs.com/)

```bash
# mac
sudo port install git-lfs

# ubuntu
sudo apt install git-lfs
```

verify that the installation was successful:

```bash
git lfs install

Updated Git hooks.
Git LFS initialized.
```

### Download a repo

- [antimatter15/alpaca.cpp](https://github.com/antimatter15/alpaca.cpp)
- [Sosaka/Alpaca-native-4bit-ggml](https://huggingface.co/Sosaka/Alpaca-native-4bit-ggml)
- [tloen/alpaca-lora-7b](https://huggingface.co/tloen/alpaca-lora-7b)
- [beomi/KoAlpaca](https://huggingface.co/beomi/KoAlpaca)
- [PotatoSpudowski/fastLLaMa](https://github.com/PotatoSpudowski/fastLLaMa)

```bash
git clone https://github.com/antimatter15/alpaca.cpp.git
git clone https://huggingface.co/Sosaka/Alpaca-native-4bit-ggml
git clone https://github.com/PotatoSpudowski/fastLLaMa
```

---

## Code

- [hello.ipynb](hello.ipynb)

### Start Jupyter

```bash
jupyter lab
```
