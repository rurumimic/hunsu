# Setup Copilot

## Getting Started in Editor

- VSCode Extension: [GitHub Copilot](https://marketplace.visualstudio.com/items?itemName=GitHub.copilot)
- Neovim Plugin: [copilot.vim](https://github.com/github/copilot.vim)
- IntelliJ Plugin: [GitHub Copilot](https://plugins.jetbrains.com/plugin/17718-github-copilot)

### SpaceVim

#### Install Node.js v16

- rurumimic/supply/[node](https://github.com/rurumimic/supply/blob/master/languages/node.md)

setup yarn.js:

```bash
nvm ls
nvm install lts/gallium
nvm alias default lts/gallium
corepack enable
```

check versions:

```bash
node -v # v16.19.1
yarn -v # 1.22.19
```

install neovim package:

```bash
npm install -g neovim
yarn global add neovim
```

#### CheckHealth SpaceVim

```bash
vim
:CheckHealth
```

```markdown
## Node.js provider (optional)
  - INFO: Node.js: v16.19.1
  - INFO: Nvim node.js host: ~/.config/yarn/global//node_modules/neovim/bin/cli.js
  - OK: Latest "neovim" npm/yarn/pnpm package is installed: 4.10.1
```

#### Add Copilot Plugin

```bash
vi ~/.SpaceVim.d/init.toml
```

```toml
[[custom_plugins]]
repo = 'github/copilot.vim'
merged = false
```

#### Enable Copilot

```bash
vim
:Copilot setup
:Copilot enable
```

```bash
:help copilot
```

```bash
:Copilot panel
```
