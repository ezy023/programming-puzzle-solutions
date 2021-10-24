use std::fmt::format;
use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn get_hint(secret: String, guess: String) -> String {
        let mut secret_chars : HashMap<u8, Vec<usize>> = HashMap::new();
        let mut guess_chars : HashMap<u8, Vec<usize>> = HashMap::new();
        for (pos, b) in secret.as_str().bytes().enumerate() {
            match secret_chars.get_mut(&b) {
                Some(v) => v.push(pos),
                None => {
                    secret_chars.insert(b, vec![pos]);
                },
            }
        }

        for (pos, b) in guess.as_str().bytes().enumerate() {
            match guess_chars.get_mut(&b) {
                Some(v) => v.push(pos),
                None => {
                    guess_chars.insert(b, vec![pos]);
                },
            }
        }

        let mut a: i32 = 0;
        let mut b: i32 = 0;

        for (k, v) in guess_chars.iter() {
            if let Some(spos) = secret_chars.get(k) {
                let mut count = 0;
                for gpos in v {
                    if spos.contains(gpos) {
                        count += 1;
                    }
                }
                a += count as i32;
                if v.len() >= spos.len() {
                    b += spos.len() as i32 - count as i32;
                } else {
                    b += v.len() as i32 - count as i32;
                }
            }
        }


        format(format_args!("{}A{}B", a, b))
    }
}



mod test {
    use super::*;

    #[test]
    fn test_get_hint() {
        assert_eq!(String::from("1A3B"), Solution::get_hint(String::from("1807"), String::from("7810")));
        assert_eq!(String::from("1A1B"), Solution::get_hint(String::from("1123"), String::from("0111")));
        assert_eq!(String::from("0A0B"), Solution::get_hint(String::from("1"), String::from("0")));
        assert_eq!(String::from("1A0B"), Solution::get_hint(String::from("1"), String::from("1")));
        assert_eq!(String::from("1A0B"), Solution::get_hint(String::from("11"), String::from("10")));
        assert_eq!(String::from("0A1B"), Solution::get_hint(String::from("1122"), String::from("0001")));
    }
}
