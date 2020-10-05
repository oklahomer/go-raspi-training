# Introduction
This repository contains some example codes and instructions to use `periph.io` library for Raspberry Pi interaction.
Six years ago, the author used Python to work with Raspberry Pi and it worked pretty well.
As a matter of fact, reading [picamera](https://github.com/waveform80/picamera)'s code and [trying out its latest feature](https://blog.oklahome.net/2014/11/trying-out-picameras-overlay-function.html) was a lot of fun.

Lately, the author prefers to write Golang for private project in favor of type-safety and compiled single binary.
This project aims to prep the author to use Golang for Raspberry Pi interaction.

`Periph.io`'s library is used as it has no C dependency, and hence cross-compilation is easier.

# Examples
- [Blink LED](https://github.com/oklahomer/go-raspi-training/tree/master/example/001-blink-led)
- [Read tactile switch state](https://github.com/oklahomer/go-raspi-training/tree/master/example/002-read-tactile-switch)
- [Measure temperature with ADT7410](https://github.com/oklahomer/go-raspi-training/tree/master/example/003-i2c/adt7410)

# Cross Compile
Take a look at `uname -a` output to decide options to cross-compile Go code.
```
pi@raspberrypi:~ $ uname -a
Linux raspberrypi 5.4.35-v7l+ #1314 SMP Fri May 1 17:47:34 BST 2020 armv7l GNU/Linux
pi@raspberrypi:~ $ cat /proc/cpuinfo
processor	: 0
model name	: ARMv7 Processor rev 3 (v7l)
BogoMIPS	: 135.00
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd08
CPU revision	: 3

processor	: 1
model name	: ARMv7 Processor rev 3 (v7l)
BogoMIPS	: 135.00
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd08
CPU revision	: 3

processor	: 2
model name	: ARMv7 Processor rev 3 (v7l)
BogoMIPS	: 135.00
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd08
CPU revision	: 3

processor	: 3
model name	: ARMv7 Processor rev 3 (v7l)
BogoMIPS	: 135.00
Features	: half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm crc32 
CPU implementer	: 0x41
CPU architecture: 7
CPU variant	: 0x0
CPU part	: 0xd08
CPU revision	: 3

Hardware	: BCM2711
Revision	: b03112
Serial		: 100000002d01a8c3
Model		: Raspberry Pi 4 Model B Rev 1.2
```

The above output indicates the Raspberry Pi's model is `ARMv7`.
For such model, The official [GoArm](https://github.com/golang/go/wiki/GoArm) wiki tells to use the GOARM value of `7` and GOARCH value of `arm` for build.
```
$ env GOOS=linux GOARCH=arm GOARM=7 go build ./example/XXX-foo-bar/main.go
```

To run this, `scp` the resulting executable to the Raspberry Pi as below.
Remember the IP address varies.
```
scp main pi@192.168.2.111:~/main
pi@192.168.2.111's password: 
main
```
