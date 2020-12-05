extern crate libc; // Cの機能があるパッケージ https://github.com/rust-lang/libc
use std::ffi::{CStr, CString};

#[no_mangle] // Rustコンパイラが関数名を変えたり削除しないように必要
pub extern "C" fn rustaceanize(name: *const libc::c_char) -> *const libc::c_char {
    let buf_name = unsafe { CStr::from_ptr(name).to_bytes() }; // CStrはlibcの型を利用してCStr型を作っている
    let mut str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("Rustaceanizing \"{}\"", str_name);
    let r_string: &str = " (V)[0-0](V)";
    str_name.push_str(r_string);
    CString::new(str_name).unwrap().into_raw()
}
