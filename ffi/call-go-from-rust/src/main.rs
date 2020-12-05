use std::ffi::{CStr, CString};
use std::os::raw::c_char;

extern "C" {
    fn Gophernize(name: GoString) -> *const c_char;
}

/// See [here](http://blog.ralch.com/tutorial/golang-sharing-libraries/) for `GoString` struct layout
#[repr(C)]
struct GoString {
    a: *const c_char,
    b: i64,
}

fn main() {
    let s = CString::new("I'm a Rustacean").expect("CString::new failed");
    let ptr = s.as_ptr();
    let input = GoString {
        a: ptr,
        b: s.as_bytes().len() as i64,
    };

    let result = unsafe { Gophernize(input) };
    let c_str = unsafe { CStr::from_ptr(result) };
    let output = c_str.to_str().expect("to_str failed");
    println!("{}", output);
}
