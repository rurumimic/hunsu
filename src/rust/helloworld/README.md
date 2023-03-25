# Hello World with Copilot

## Create a project

```bash
cargo new helloworld
```

### with Copilot

```bash
vi src/main.rs
```

```rust
// Take the user's name and say hello
```

1. `:Copilot panel`
2. Select a suggestion
3. `Enter`

#### Codes

```rust
use std::io;

fn main() {
    let name = get_name();
    println!("Hello, {}!", name);
}

// Path: helloworld/src/main.rs
// Get the user's name from stdin
fn get_name() -> String {
    let mut name = String::new();
    println!("What is your name?");
    let stdin = io::stdin();
    stdin
        .read_line(&mut name)
        .ok()
        .expect("Failed to read line");
    name
}
```

#### Run Copilot's Code

```bash
cargo run
```

```bash
What is your name?

GitHub Copilot
```

```bash
What is your name?
GitHub Copilot
Hello, GitHub Copilot
!
```

The newline was not handled correctly  
!
