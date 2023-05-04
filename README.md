# 훈수

written by GitHub Copilot and me.

---

## Baduk

![78](/images/78-sabaki-leelazero.png)

- Game 4: Lee Sedol vs Alphago

### 신진서

- Shin Jinseo's [Rating](https://www.goratings.org/en/players/1313.html)
- 🗞️ ["AI is for reference only, I make my own moves"](https://www.hankyung.com/life/article/2022030167811)
- 🗞️ ["AI stunts, heart pounding"](https://www.mk.co.kr/news/society/9775037)

> Thanks to AI, my understanding deepens as I keep playing one style. It's relatively easy to predict against other players who study AI.

> During a big game, I sometimes ponder, "If I were an AI, where would I place my move?". However, it's crucial to find my own best move. I think I should learn from AI and develop my unique strengths.

> It's undeniable that AI has provided immense help in advancing my Go skills. However, no matter how perfect the AI is, I believe there is still room for human spirit and competition to shine through. While playing Go, there are moves that the AI does not recommend, yet have a higher probability than what AI suggests. These moves may only come up once in a game, but I believe they represent human potential.

### AI

- KataGo: [github](https://github.com/lightvector/KataGo): GTP engine and self-play learning in Go
- Leela Zero: [github](https://github.com/leela-zero/leela-zero): Go engine with no human-provided knowledge, modeled after the AlphaGo Zero paper.

### GUI

- Lizzie: [github](https://github.com/featurecat/lizzie), [releases](https://github.com/featurecat/lizzie/releases)
  - yzyray/lizzieyzy: [github](https://github.com/yzyray/lizzieyzy), [english](https://github.com/yzyray/lizzieyzy/blob/main/README_EN.md), [manual.pdf](https://github.com/yzyray/lizzieyzy/blob/main/readme_en.pdf), [releases](https://github.com/yzyray/lizzieyzy/releases)
- Sabaki: [github](https://github.com/SabakiHQ/Sabaki)

### Play a game

1. [Play a game](docs/baduk.md)

---

## Copilot

- [GitHub Copilot](https://copilot.github.com/) is a new AI-powered code completion tool that helps you write code faster and more efficiently.
  - [docs](https://docs.github.com/en/copilot)

### with Copilot

1. [Setup Copilot](docs/setup.md): Vscode, Neovim, SpaceVim, IntelliJ
   - [SpaceVim](docs/setup.md#spacevim): [Add Copilot Plugin](docs/setup.md#add-copilot-plugin)

---

## OpenAI

- [OpenAI](https://openai.com/)
  - [platform](https://platform.openai.com)
    - [docs](https://platform.openai.com/docs)
- Code
  - [cookbook](https://github.com/openai/openai-cookbook)

### ChatGPT

#### Awesome

- [propmts.chat](https://prompts.chat): A collection of prompts for the ChatGPT model.
- [ChainBrain AI](https://www.chainbrainai.com/): Advanced ChatGPT prompts.

### Whisper

- Whisper: [github](https://github.com/openai/whisper)

---

## Microsoft

### JARVIS

- JARVIS
  - [github](https://github.com/microsoft/JARVIS)
  - [paper](https://arxiv.org/abs/2303.17580): Hugging GPT

---

## And also

### Marvin

- [Marvin](https://www.askmarvin.ai/)
  - [github](https://github.com/PrefectHQ/marvin): A batteries-included library for building AI-powered software
  - [install](https://www.askmarvin.ai/getting_started/installation/)

### Bloomberg

- BloombergGPT
  - 50-billion parameter LLM tuned for finance
  - [announcements](https://www.bloomberg.com/company/press/bloomberggpt-50-billion-parameter-llm-tuned-finance/)
  - [paper](https://arxiv.org/abs/2303.17564): [summary](papers/bloomberg/bloomberggpt.md)

### LMQL

- [LMQL](https://lmql.ai/)
  - [github](https://github.com/eth-sri/lmql)

---

## Code

- [Hello World](src/rust/helloworld/README.md) with Copilot
  - [helloworld/src/main.rs](src/rust/helloworld/src/main.rs)
- OpenAI
  - [Hello](src/openai/go/hello/README.md)
- Marvin
  - [start](src/marvin/start/README.md): install, start
    - List of N Fruits: [fruits.py](src/marvin/start/fruits.py), [prompt](src/marvin/start/prompt/fruits.md)
    - Classify Tweets: [tweets.py](src/marvin/start/tweets.py), [prompt](src/marvin/start/prompt/tweets.md)
- LMQL
  - [hello](src/lmql/start/README.md): install, start
    - Hello: [hello.py](src/lmql/start/hello.py)
- My LLM
  - [alpaca](src/myllm/alpaca/README.md)
