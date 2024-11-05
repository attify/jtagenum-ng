use rppal::gpio::{Bias, Gpio, IoPin, Level};
use std::collections::HashMap;

pub struct RppalDriver {
    gpio: Box<Gpio>,
    alloc_pins: HashMap<u8, IoPin>,
}

impl RppalDriver {
    fn init() -> RppalDriver {
        let gpio = Box::new(Gpio::new().unwrap());
        RppalDriver {
            gpio: gpio,
            alloc_pins: Default::default(),
        }
    }

    fn get_alloc_pin(&mut self, gpio_pin: u8) -> &mut IoPin {
        if !self.alloc_pins.contains_key(&gpio_pin) {
            let new_pin = self
                .gpio
                .get(gpio_pin)
                .unwrap()
                .into_io(rppal::gpio::Mode::Input);
            self.alloc_pins.insert(gpio_pin, new_pin);
        }
        self.alloc_pins.get_mut(&gpio_pin).unwrap()
    }

    fn pin_input(&mut self, gpio_pin: u8) {
        let pin = self.get_alloc_pin(gpio_pin);
        pin.set_mode(rppal::gpio::Mode::Input);
    }

    fn pin_output(&mut self, gpio_pin: u8) {
        let pin = self.get_alloc_pin(gpio_pin);
        pin.set_mode(rppal::gpio::Mode::Output);
    }

    fn pin_write(&mut self, gpio_pin: u8, value: u8) {
        let pin = self.get_alloc_pin(gpio_pin);
        pin.write(if value == 0 { Level::Low } else { Level::High });
    }

    fn pin_read(&mut self, gpio_pin: u8) -> u8 {
        let pin = self.get_alloc_pin(gpio_pin);
        if pin.read() == Level::High {
            1
        } else {
            0
        }
    }

    fn pin_pull_up(&mut self, gpio_pin: u8) {
        let pin = self.get_alloc_pin(gpio_pin);
        pin.set_bias(Bias::PullUp);
    }

    fn pin_pull_off(&mut self, gpio_pin: u8) {
        let pin = self.get_alloc_pin(gpio_pin);
        pin.set_bias(Bias::Off);
    }
}

#[no_mangle]
pub extern "C" fn init_driver() -> *mut RppalDriver {
    let obj = Box::new(RppalDriver::init());
    Box::into_raw(obj)
}

#[no_mangle]
pub extern "C" fn close_driver(ptr: *mut RppalDriver) {
    if !ptr.is_null() {
        unsafe {
            drop(Box::from_raw(ptr));
        }
    }
}

#[no_mangle]
pub extern "C" fn pin_write(ptr: *mut RppalDriver, gpio_pin: u8, value: u8) {
    if !ptr.is_null() {
        unsafe {
            (*ptr).pin_write(gpio_pin, value);
        }
    }
}

#[no_mangle]
pub extern "C" fn pin_read(ptr: *mut RppalDriver, gpio_pin: u8) -> u8 {
    if !ptr.is_null() {
        unsafe { (*ptr).pin_read(gpio_pin) }
    } else {
        0
    }
}

#[no_mangle]
pub extern "C" fn pin_output(ptr: *mut RppalDriver, gpio_pin: u8) {
    if !ptr.is_null() {
        unsafe {
            (*ptr).pin_output(gpio_pin);
        }
    }
}

#[no_mangle]
pub extern "C" fn pin_input(ptr: *mut RppalDriver, gpio_pin: u8) {
    if !ptr.is_null() {
        unsafe {
            (*ptr).pin_input(gpio_pin);
        }
    }
}

#[no_mangle]
pub extern "C" fn pin_pull_up(ptr: *mut RppalDriver, gpio_pin: u8) {
    if !ptr.is_null() {
        unsafe {
            (*ptr).pin_pull_up(gpio_pin);
        }
    }
}

#[no_mangle]
pub extern "C" fn pin_pull_off(ptr: *mut RppalDriver, gpio_pin: u8) {
    if !ptr.is_null() {
        unsafe {
            (*ptr).pin_pull_off(gpio_pin);
        }
    }
}
