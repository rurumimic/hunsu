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

---

## Code

- [alpaca_fastllama.ipynb](alpaca_fastllama.ipynb)
- [koalpaca.ipynb](koalpaca.ipynb)

### Start Jupyter

```bash
jupyter lab
```
