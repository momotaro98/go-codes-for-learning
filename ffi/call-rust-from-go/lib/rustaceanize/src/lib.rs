extern crate libc;
use std::ffi::{CStr, CString};

#[no_mangle] // no_mangle はRustコンパイラが関数名を変えたり削除しないように必要
pub extern "C" fn rustaceanize(name: *const libc::c_char) -> *const libc::c_char {
    let cstr_name = unsafe { CStr::from_ptr(name) };
    let mut str_name = cstr_name.to_str().unwrap().to_string();
    println!("Rustaceanizing \"{}\"", str_name);
    let r_string: &str = " (V)[0-0](V)";
    str_name.push_str(r_string);
    CString::new(str_name).unwrap().into_raw()
}
