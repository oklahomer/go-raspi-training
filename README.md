# Introduction
This repository contains some example codes and instructions to use `periph.io` library for Raspberry Pi interaction.
Six years ago, the author used Python to work with Raspberry Pi and it worked pretty well.
As a matter of fact, reading [picamera](https://github.com/waveform80/picamera)'s code and [trying out its latest feature](https://blog.oklahome.net/2014/11/trying-out-picameras-overlay-function.html) was a lot of fun.

Lately, the author prefers to write Golang for private project in favor of type-safety and compiled single binary.
This project aims to prep the author to use Golang for Raspberry Pi interaction.

`Periph.io`'s library is used as it has no C dependency, and hence cross-compilation is easier.

# Examples
- [Blink LED manually and with periph.io](https://github.com/oklahomer/go-raspi-training/tree/master/example/001-blink-led)
