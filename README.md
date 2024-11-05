# Jtagenum-NG

> This project is aimed to find which pins exposed by the target device are JTAG pins. It does so by enumerating throughout the provided pins set and trying to abuse some JTAG features, such as BYPASS and IDCODE registers.

This project is targeted at Raspberry Pi 5 (with a 64bit Linux OS installed) and hence binaries have only been provided for that platform.

## Usage

Precompiled binaries are provided in Releases. Otherwise for compiling from source, refer to the [build workflow](.github/workflows/build.yaml)

Please refer to the original project for more details: https://github.com/gremwell/go-jtagenum?tab=readme-ov-file#usage

## Changes compared to Go-JTAGenum

- Removed the go-rpio driver as it doesn't support RPi 5.
- Removed the libgpiod driver. This driver does work with Pi 5 but requires the `libgpiod-dev` package to be installed on the Pi. Also linking against the libraries can be tricky specifically when cross-compiling on a x86_64 host.
- Added a driver based on [go-gpiocdev](https://github.com/warthog618/go-gpiocdev). This is a pure Go library and uses the Linux GPIO character device via `ioctls` to access the Pi's pins.
- Added a driver based on [rppal](https://github.com/golemparts/rppal) library. This is a Rust library that provides access to the Raspberry Pi GPIO ports and peripherals. To allow the Go code to interact with Rust a separate driver has been built in *driver_rppal* directory. Rppal bypasses the kernel and uses memory mapped IO (MMIO) to interact with the GPIO pins.

## Credits

Go-Jtagenum: https://github.com/gremwell/go-jtagenum

## License

GNU General Public License v3.0