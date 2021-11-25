struct SnapshotArray {
    snapshot_id: i32,
    // list of list of tuples. The top index represents the element. The list at that index represents changes to the element with 'set'. The tuple is (value, snapshot_id) where
    // snapshot_id is the snapshot_id at the time of the 'set' operation.
    store: Vec<Vec<(i32,i32)>>,
}

impl SnapshotArray {
    fn new(length: i32) -> Self {
        let mut store = Vec::new();
        for _i in 0..length {
            store.push(vec![(0,0)]);
        }
        SnapshotArray{snapshot_id: 0, store}
    }

    fn set(&mut self, index: i32, val: i32) {
        self.store[index as usize].push((val, self.snapshot_id));
    }

    fn snap(&mut self) -> i32 {
        self.snapshot_id += 1;
        self.snapshot_id - 1
    }

    fn get(&self, index: i32, snap_id: i32) -> i32 {
        let element_list = &self.store[index as usize];
        for i in (0..element_list.len()).rev() {
            let element = element_list[i];
            if element.1 <= snap_id {
                return element.0
            }
        }
        0
    }
}



fn runner(commands: Vec<String>, arguments: Vec<Vec<i32>>) -> Vec<Option<i32>> {
    let mut ziperator = commands.iter().zip(arguments.iter());

    let mut snapshot =  match ziperator.next() {
        Some((command, arg)) => {
            if command != "SnapshotArray" {
                panic!("First command must be to create a new SnapshotArray");
            }
            if arg.len() != 1 {
                panic!("SnapshotArray::new accepts only a single argument, given {:?}", arg);
            }
            SnapshotArray::new(arg[0])
        },
        None => panic!("Expected non-empty input args"),
    };

    let mut output = vec![None; 1];

    for (command, args) in ziperator {
        match command.as_str() {
            "set" => {
                // could check arg lengths, but would panic anyways so let it panic on the access
                snapshot.set(args[0], args[1]);
                output.push(None);
            },
            "snap" => {
                output.push(Some(snapshot.snap()));
            },
            "get" => {
                let index = args[0];
                let snapshot_id = args[1];
                output.push(Some(snapshot.get(index, snapshot_id)));
            },
            _ => panic!("Unexpected command {}", command),
        }
    }

    output
}

#[cfg(test)]
mod test {
    use crate::*;

    #[test]
    fn test_example_one() {
        let commands = vec![
            String::from("SnapshotArray"),
            String::from("set"),
            String::from("snap"),
            String::from("set"),
            String::from("get")
        ];
        let args = vec![
            vec![3],
            vec![0,5],
            vec![],
            vec![0,6],
            vec![0,0],
        ];

        let expected = [None, None, Some(0), None, Some(5)];
        let gots = runner(commands, args);
        for (exp, got) in expected.iter().zip(gots.iter()) {
            assert_eq!(exp, got);
        }
    }

    #[test]
    fn test_custom_one() {
        let commands = vec![
            String::from("SnapshotArray"),
            String::from("snap"),
            String::from("snap"),
            String::from("snap"),
            String::from("snap"),
            String::from("snap"),
            String::from("snap"),
            String::from("get"),
        ];
        let args = vec![
            vec![1000],
            vec![],
            vec![],
            vec![],
            vec![],
            vec![],
            vec![],
            vec![500, 4], // element and index 500 of snapshot 4
        ];

        let expected = [None, Some(0), Some(1), Some(2), Some(3), Some(4), Some(5), Some(0)];
        let gots = runner(commands, args);
        for (exp, got) in expected.iter().zip(gots.iter()) {
            assert_eq!(exp, got);
        }
    }
}
