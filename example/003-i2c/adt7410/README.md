# Measure temperature with ADT7410

## Diagram

![](https://raw.githubusercontent.com/oklahomer/go-raspi-training/master/example/003-i2c/adt7410/img/diagram.png)

## Manual procedure
Before start working with ADT7410 temperature sensor, a developer has to enable I²C interface with `raspi-config`.
```
pi@raspberrypi:~ $ sudo raspi-config
```

Under "Interfacing Options" > "I2C," there is a menu to enable I²C.

![](https://raw.githubusercontent.com/oklahomer/go-raspi-training/master/example/003-i2c/adt7410/img/raspi-config.png)

Then execute the below command and install some useful tools:
```
pi@raspberrypi:~ $ sudo apt-get install -y i2c-tools
```

Now `i2cdetect` is installed.
List all installed busses as below:
```
pi@raspberrypi:~ $ i2cdetect -l
i2c-1	i2c       	bcm2835 (i2c@7e804000)          	I2C adapter
```

Above output indicates that an I²C bus with number 1 is available.
`i2cdetect -y 1` will list all detected devices for i2c-**1**.
After wiring, designated address for the device can be shown as part of this result.

```
pi@raspberrypi:~ $ i2cdetect -y 1
     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
00:          -- -- -- -- -- -- -- -- -- -- -- -- --
10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
40: -- -- -- -- -- -- -- -- 48 -- -- -- -- -- -- --
50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
70: -- -- -- -- -- -- -- --
```

The result says the address is 48.
Note the output shows the address in a hexadecimal form so this has to be declared as `0x48` in application code.

## Go implementation with periph.io
Take a look at [Cross Compile](https://github.com/oklahomer/go-raspi-training/#cross-compile) to cross-compile and transfer the example code.
The corresponding Go code is located at [/example/003-i2c/adt7410/main.go](https://github.com/oklahomer/go-raspi-training/blob/master/example/003-i2c/adt7410/main.go).

Login and execute the transferred executable.
```
pi@raspberrypi:~ $ ./main 
2020/10/03 12:04:19 Start
2020/10/03 12:04:22 Current temperature: 27.437500°C (81.387497°F)
2020/10/03 12:04:25 Current temperature: 27.437500°C (81.387497°F)
2020/10/03 12:04:28 Current temperature: 27.437500°C (81.387497°F)
^C2020/10/03 12:04:32 Stopped
2020/10/03 12:04:32 Closed I²C device
```