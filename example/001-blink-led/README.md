# Blink LED with periph.io
## Diagram
![](https://raw.githubusercontent.com/oklahomer/go-raspi-training/master/example/001-blink-led/img/diagram.png)

## Manual Procedure
Before programmatically blink the LED light with higher level interface, this section introduces how to manually 
interact with the LED light by accessing file system.

At the beginning, files and directories under `/sys/class/gpio` look somewhat like below:
```
pi@raspberrypi:~ $ ls /sys/class/gpio/
export  gpiochip0  gpiochip504  unexport
```

To interact with the LED light via GPIO 23, the pin needs to be "exported."
```
pi@raspberrypi:~ $ echo 23 > /sys/class/gpio/export
pi@raspberrypi:~ $ ls /sys/class/gpio/
export  gpio23  gpiochip0  gpiochip504  unexport
```

Its direction is `in` at this point.
```
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/direction 
in
```

Override it with `out` to make the corresponding pin an output pin.
```
pi@raspberrypi:~ $ echo out > /sys/class/gpio/gpio23/direction 
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/direction 
out
```

To change the voltage level of the output pin and blink, override the file content.
Set `1` to light up; set `0` to dim.
```
pi@raspberrypi:~ $ echo 1 > /sys/class/gpio/gpio23/value 
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/value 
1
pi@raspberrypi:~ $ echo 0 > /sys/class/gpio/gpio23/value 
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/value 
0
```

When everything is done, "unexport" the exported pin to finish the ongoing interaction.
```
pi@raspberrypi:~ $ echo 23 > /sys/class/gpio/unexport
pi@raspberrypi:~ $ ls /sys/class/gpio/
export  gpiochip0  gpiochip504  unexport
```

## Go implementation with periph.io
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
$ env GOOS=linux GOARCH=arm GOARM=7 go build ./main.go
```

`SCP` the resulting executable to Raspberry Pi as below. Remember the IP address varies.
```
scp main pi@192.168.2.111:~/main
pi@192.168.2.111's password: 
main
```

Login and execute the transferred executable.
By its execution, files and directories under `/sys/class/gpio/` are implicitly initialized.
The value of `/sys/class/gpio/gpio23/value` changes every 3 seconds until `main.go` stops.
```
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/value 
1
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/value 
0
```
