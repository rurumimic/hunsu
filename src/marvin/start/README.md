# Marvin

## Setup

### Set Python

- rurumimic: [python](https://github.com/rurumimic/supply/blob/master/languages/python.md)

```bash
pyenv virtualenv 3.11.1 marvin
pyenv activate marvin
pip install --upgrade pip setuptools wheel
pip install black
```

### Install Marvin

```bash
pip install marvin
```

### OpenAI API Key

OpenAI: [API Keys](https://platform.openai.com/account/api-keys)

1. Create a new secret key.
2. Save a key: `.env`

```bash
marvin setup-openai
```

```bash
[04/05/23 14:54:15] INFO     marvin.marvin: Using OpenAI model "gpt-3.5-turbo"
Enter your OpenAI API key to save it to /Users/keanu/.marvin/.env, or hit enter to unset:
```

#### Marvin Key

```bash
# $HOME/.marvin/.env
MARVIN_OPENAI_API_KEY=''
```

---

## Code

### List of N Fruits

- List of N Fruits: [fruits.py](fruits.py)
- prompt: [history](prompt/fruits.md)

```bash
python fruits.py

[04/05/23 15:10:08] INFO     marvin.marvin: Using OpenAI model "gpt-3.5-turbo"
과일 이름 3개 목록을 만들어
['apple', 'banana', 'orange']
```

### Classify Tweets

- Classify Tweets: [tweets.py](tweets.py)
- prompt: [history](prompt/tweets.md)

```bash
python tweets.py

[04/05/23 15:10:14] INFO     marvin.marvin: Using OpenAI model "gpt-3.5-turbo"
트윗을 분류해
False: @elonmusk: I wish the media would stop flattering me all time … it’s a bit much guys ☺️
True: @realDonaldTrump: Please support our Capitol Police and Law Enforcement. They are truly on the side of our Country. Stay peaceful!
True: @Sonny7: A legend of Spurs and a legend of football. Congratulations on an amazing career @GarethBale11 !! Wishing you the best luck in your next chapter mate 😁🤍
```
