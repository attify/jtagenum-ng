package main

//#cgo LDFLAGS: -L../driver_rppal/target/aarch64-unknown-linux-gnu/release -ldriver_rppal
//#include "../driver_rppal/bindings.h"
import "C"

type JtagPinDriverRppal struct {
	rppalDriver *C.RppalDriver
}

func (d *JtagPinDriverRppal) initDriver() {
	d.rppalDriver = C.init_driver()
}

func (d *JtagPinDriverRppal) closeDriver() {
	C.close_driver(d.rppalDriver)
}

func (d *JtagPinDriverRppal) pinWrite(pin JtagPin, state JtagPinState) {
	C.pin_write(d.rppalDriver, C.uchar(pin), C.uchar(state))
}

func (d *JtagPinDriverRppal) pinRead(pin JtagPin) JtagPinState {
	return JtagPinState(C.pin_read(d.rppalDriver, C.uchar(pin)))
}

func (d *JtagPinDriverRppal) pinOutput(pin JtagPin) {
	C.pin_output(d.rppalDriver, C.uchar(pin))
}

func (d *JtagPinDriverRppal) pinInput(pin JtagPin) {
	C.pin_input(d.rppalDriver, C.uchar(pin))
}

func (d *JtagPinDriverRppal) pinPullUp(pin JtagPin) {
	C.pin_pull_up(d.rppalDriver, C.uchar(pin))
}

func (d *JtagPinDriverRppal) pinPullOff(pin JtagPin) {
	C.pin_pull_off(d.rppalDriver, C.uchar(pin))
}