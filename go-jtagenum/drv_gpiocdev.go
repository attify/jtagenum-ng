package main

import (
	"fmt"
	"github.com/warthog618/go-gpiocdev"
)

type JtagPinDriverGpioCDev struct {
	GpioChipIdx uint
	chip      *gpiocdev.Chip
	lines    map[JtagPin]*gpiocdev.Line
}

func (d *JtagPinDriverGpioCDev) initDriver() {
	var err error
	d.chip, err = gpiocdev.NewChip(fmt.Sprintf("gpiochip%d", d.GpioChipIdx))
	if err != nil {
		panic(fmt.Sprintf("can't open gpio chip #%d", d.GpioChipIdx))
	}
	d.lines = make(map[JtagPin]*gpiocdev.Line, 0)
}

func (d *JtagPinDriverGpioCDev) closeDriver() {
	for _, line := range d.lines {
		line.Close()
	}
	d.chip.Close()
}

func (d *JtagPinDriverGpioCDev) getAllocLine(pin JtagPin) *gpiocdev.Line {
	var l *gpiocdev.Line
	var ok bool
	l, ok = d.lines[pin]
	if !ok {
		var err error
		l, err = d.chip.RequestLine(int(pin))
		if err != nil {
			panic(fmt.Sprintf("can't reserve pin #%d", pin))
		}
		d.lines[pin] = l
	}
	return l
}

func (d *JtagPinDriverGpioCDev) pinWrite(pin JtagPin, state JtagPinState) {
	d.getAllocLine(pin).SetValue(int(state))
}

func (d *JtagPinDriverGpioCDev) pinRead(pin JtagPin) JtagPinState {
	v, err := d.getAllocLine(pin).Value()
	if err != nil {
		panic(fmt.Sprintf("can't get pin #%d value", pin))
	}
	return JtagPinState(v)
}

func (d *JtagPinDriverGpioCDev) pinOutput(pin JtagPin) {
	l, ok := d.lines[pin]
	if !ok {
		l = d.getAllocLine(pin)
	}
	l.Reconfigure(gpiocdev.AsOutput(1))
}

func (d *JtagPinDriverGpioCDev) pinInput(pin JtagPin) {
	l, ok := d.lines[pin]
	if !ok {
		l = d.getAllocLine(pin)
	}
	l.Reconfigure(gpiocdev.AsInput)	
}

func (d *JtagPinDriverGpioCDev) pinPullUp(pin JtagPin) {
	l, ok := d.lines[pin]
	if !ok {
		l = d.getAllocLine(pin)
	}
	l.Reconfigure(gpiocdev.WithPullUp)
}

func (d *JtagPinDriverGpioCDev) pinPullOff(pin JtagPin) {
	l, ok := d.lines[pin]
	if !ok {
		l = d.getAllocLine(pin)
	}
	l.Reconfigure(gpiocdev.WithBiasDisabled)
}
