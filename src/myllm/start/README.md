# Custom

## Setup Env

### Python Env

```bash
pyenv virtualenv 3.11.1 myllm
pyenv activate myllm
```

```bash
pip install --upgrade pip setuptools wheel
pip install black
```

### Install Library

```bash
pip install -r requirements.txt
```

### Download model repo

#### git lfs

```bash
sudo port install git-lfs
```

verify that the installation was successful:

```bash
git lfs install

Updated Git hooks.
Git LFS initialized.
```

#### download a repo

- huggingface: [beomi/KoAlpaca](https://huggingface.co/beomi/KoAlpaca)

```bash
git clone https://huggingface.co/beomi/KoAlpaca
```

---

## Code

- [hello.ipynb](hello.ipynb)
- [app.py](app.py)
