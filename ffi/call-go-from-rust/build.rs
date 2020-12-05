fn main() {
    let path = "./golib";
    let lib = "gophernize";

    println!("cargo:rustc-link-search=native={}", path);
    println!("cargo:rustc-link-lib=static={}", lib);
}
