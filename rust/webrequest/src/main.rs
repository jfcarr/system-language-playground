extern crate env_logger;
extern crate reqwest;

fn main() -> Result<(), Box<std::error::Error>> {
    env_logger::init();

    let the_url = "https://www.rust-lang.org";

    println!("GET {0}", the_url);

    let mut res = reqwest::get(the_url)?;

    println!("Status: {}", res.status());
    println!("Headers:\n{:?}", res.headers());

    // copy the response body directly to stdout
    std::io::copy(&mut res, &mut std::io::stdout())?;

    println!("\n\nDone.");

    Ok(())
}
