/* THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT */

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct RppalDriver RppalDriver;

struct RppalDriver *init_driver(void);

void close_driver(struct RppalDriver *ptr);

void pin_write(struct RppalDriver *ptr, uint8_t gpio_pin, uint8_t value);

uint8_t pin_read(struct RppalDriver *ptr, uint8_t gpio_pin);

void pin_output(struct RppalDriver *ptr, uint8_t gpio_pin);

void pin_input(struct RppalDriver *ptr, uint8_t gpio_pin);

void pin_pull_up(struct RppalDriver *ptr, uint8_t gpio_pin);

void pin_pull_off(struct RppalDriver *ptr, uint8_t gpio_pin);
