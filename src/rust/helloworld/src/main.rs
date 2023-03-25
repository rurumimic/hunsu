// Take the user's name and say hello
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
