# Build Instructions

Assuming a 64 bit OS is installed such as Ubuntu Server 64 bit (arm64)

### Prerequisites

```
sudo apt install build-essential gcc-aarch64-linux-gnu
```

Note that for arm 32bit requires  `gcc-arm-linux-gnueabi` or `gcc-arm-linux-gnueabihf`

## Rust side

One time setup
```
rustup default stable
rustup target add aarch64-unknown-linux-gnu
```

### Build

```
cargo build --release --target=aarch64-unknown-linux-gnu
```
Omit `--release` for debug build. This creates both *bindings.h* in the *driver_rppal* directory as well as the compiled static library *driver_librppal.a*.

## Go side

For 64-bit OS (always hard float)
```
sudo apt install gcc-aarch64-linux-gnu
CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc  GOOS=linux GOARCH=arm64 go build -ldflags="-w -s"
```


For 32-bit OS (soft float)
```
sudo apt install gcc-arm-linux-gnueabi
CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc  GOOS=linux GOARCH=arm go build -ldflags="-w -s"
```

For 32-bit OS (hard float)
```
sudo apt install gcc-arm-linux-gnueabihf
CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm go build -ldflags="-w -s"
```