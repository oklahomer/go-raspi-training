# Blink LED with periph.io
## Diagram
![](https://raw.githubusercontent.com/oklahomer/go-raspi-training/master/example/001-blink-led/img/diagram.png)

## Manual procedure
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
Take a look at [Read tactile switch state](https://github.com/oklahomer/go-raspi-training/#cross-compile) to cross-compile and transfer the example code.

Login and execute the transferred executable.
By its execution, files and directories under `/sys/class/gpio/` are implicitly initialized.
The value of `/sys/class/gpio/gpio23/value` changes every 3 seconds until `main.go` stops.
```
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/value 
1
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio23/value 
0
```
