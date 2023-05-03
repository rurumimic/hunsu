# Play a game

## Open Sources

### AI

- KataGo: [github](https://github.com/lightvector/KataGo): GTP engine and self-play learning in Go
- Leela Zero: [github](https://github.com/leela-zero/leela-zero): Go engine with no human-provided knowledge, modeled after the AlphaGo Zero paper.

### GUI

- Lizzie: [github](https://github.com/featurecat/lizzie), [releases](https://github.com/featurecat/lizzie/releases)
- yzyray/lizzieyzy: [github](https://github.com/yzyray/lizzieyzy), [english](https://github.com/yzyray/lizzieyzy/blob/main/README_EN.md), [manual.pdf](https://github.com/yzyray/lizzieyzy/blob/main/readme_en.pdf), [releases](https://github.com/yzyray/lizzieyzy/releases)

---

## Setup

### Java

- JDK 17: [Install](https://github.com/rurumimic/supply/blob/master/languages/java.md)

### Download KataGo

- KataGo: [releases](https://github.com/lightvector/KataGo/releases)

### Download or Build Leela Zero

- Leela Zero: [releases](https://github.com/leela-zero/leela-zero/releases)

#### Build Leela Zero on macOS

```bash
sudo port install boost cmake zlib

git clone https://github.com/leela-zero/leela-zero
cd leela-zero
git submodule update --init --recursive

mkdir build && cd build

cmake ..
cmake --build .
```

### Download Lizzie

- Lizzie: [releases](https://github.com/featurecat/lizzie/releases)

---

## Run Lizzie

```bash
java -jar lizzie.jar
```

---

## Lee Sedol vs Alphago

### Game 4

- Game 4: [20160313-LeeSedol-Alphago.sgf](/games/20160313-LeeSedol-Alphago.sgf)

![77](/images/77.png)
![78](/images/78.png)
![win](/images/win.png)
