extern crate sysinfo;

use sysinfo::SystemExt;

fn main() {
    let mut system = sysinfo::System::new();

    system.refresh_all();

	/*
    for process in system.get_process_list() {
        println!("{:?}", process);
    }
	*/

	for (key,value ) in system.get_process_list() {
		 println!("{:?}", key);
		 println!("{:?}", value);
	}
}
