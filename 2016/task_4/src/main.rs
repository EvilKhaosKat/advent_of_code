use std::fs::File;
use std::io::{BufRead, BufReader};
use std::collections::HashMap;

fn main() {
    let input = get_input();

    println!("{}", solve_first(input));
}

enum Checksum {
    Real(u16),
    Decoy,
}

fn solve_first(input: Vec<String>) -> u16 {
    let mut result = 0;

    for line in input.iter() {
        match validate_checksum(line) {
            Checksum::Real(id) => result += id,
            Checksum::Decoy => {}
        }
    }

    return result;
}

fn validate_checksum(line: &String) -> Checksum {
    let parts: Vec<&str> = line.split('-').collect();
    let suffix_index = parts.len() - 1;

    let encrypted_name = &parts[..suffix_index].concat();

    let suffix_parts: Vec<&str> = parts[suffix_index]
        .split(|delimiter| delimiter == '[' || delimiter == ']').collect();

    let id = suffix_parts[0].parse().unwrap();
    let checksum = suffix_parts[1];

    return validate_checksum_parsed(encrypted_name, id, checksum);
}

fn validate_checksum_parsed(encrypted_name: &str, id: u32, checksum: &str) -> Checksum {
    let mut letters_freq: HashMap<char, u8> = HashMap::new();

    for c in encrypted_name.chars() {
        let freq = letters_freq.entry(c).or_insert(0);
        *freq += 1;
    }

    //TODO

    return Checksum::Decoy;
}

fn get_input() -> Vec<String> {
    let f = File::open("input.txt").unwrap();
    let f = BufReader::new(f);

    let mut result = Vec::new();

    for line in f.lines() {
        result.push(line.unwrap());
    }

    return result;
}