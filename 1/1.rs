use std::fs::read_to_string;
use std::io;

fn arg() -> String {
    let v: Vec<String> = std::env::args().collect();
    v[1]
}

fn main() -> io::Result<()> {
    let s = read_to_string(arg())?;
    let i: Vec<i64> = s.lines().map(|l| l.parse().unwrap()).collect();
    println!("{:?}", i);
    Ok(())
}
