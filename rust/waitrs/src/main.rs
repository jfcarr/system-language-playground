use std::process::Command;

fn main() {
    let mut child = Command::new("sleep").arg("5").spawn().unwrap();

    println!("starting call");

    let _result = child.wait().unwrap();

    println!("reached end of main");
}
