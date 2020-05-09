# Read tactile switch state with periph.io
While [001-blink-led example](https://github.com/oklahomer/go-raspi-training/blob/master/example/001-blink-led/README.md)
demonstrates how output works with GPIO pins, this example shows how input works.

## Diagram
### Simpler circuit with built-in pull-down resister
Below depicts the simplest circuit to read the state of a tactile switch.
However, when the switch is not being pushed, this circuit is potentially unstable because the wire from GPIO #3 is not connected to 3.3v or the ground.
Such a state is described as "floating."

![](https://raw.githubusercontent.com/oklahomer/go-raspi-training/master/example/002-read-tactile-switch/img/diagram_internal_pulldown.png)

When the wire is connected to the 3.3v power supply, 3.3v comes as input, and hence the input level is "High";
when connected to the ground, 0.0v comes as input, and hence the input level is "Low."

NOTE: See [GPIO pads control](https://www.raspberrypi.org/documentation/hardware/raspberrypi/gpio/gpio_pads_control.md) for the threshold value of "High" and "Low."

Without the switch being pushed, GPIO pin is connected to neither of them.
Electrical noise may affect determining such a pin's input level.
The good thing about Raspberry Pi's GPIO pins is that pins have a built-in mechanism to handle such an unstable state.
This mechanism works as below:
- Pull-up: The pin is internally connected to the 3.3v rail having a register in-between, so the level is "pulled up" to high
- Pull-down: The pin is internally connected to the ground and the level is "pulled down" to low

This particular scenario uses internal pull-down register such as below:
```go
// Pass gpio.PullDown to activate internal pull-down register
err := PIN.In(gpio.PullDown, gpio.NoEdge)
```

### Circuit with explicit pull-down mechanism
While the above circuit uses internal pull-down mechanism to ensure the stability, below circuit explicitly connects the wire to the ground.
When the switch is not pushed, the wire from GPIO pin is connected to the ground and its level is low;
when pushed, the 3.3v power supply is connected and its level is high.


![](https://raw.githubusercontent.com/oklahomer/go-raspi-training/master/example/002-read-tactile-switch/img/diagram_external_pulldown.png)

Since pull-down is explicitly enabled, additional software configuration is not required in this case.

```go
// Pass gpio.Float not to activate internal pull-down/pull-up register
err := PIN.In(gpio.Float, gpio.NoEdge)
```

One more thing to note in this circuit is that the register between the power supply and the ground is also placed between the GPIO pin and the ground.
If the GPIO is mistakenly set for output and gives 3.3v, the register still protects the Raspberry Pi from high-intensity.

## Manual procedure

Export GPIO pin #3 and make sure the direction is "in" for input.
If the value of `direction` file is "out," then set this to "in" by `echo in > /sys/class/gpio/gpio3/direction`
```
pi@raspberrypi:~ $ ls /sys/class/gpio/
export  gpiochip0  gpiochip504  unexport
pi@raspberrypi:~ $ echo 3 > /sys/class/gpio/export
pi@raspberrypi:~ $ ls /sys/class/gpio/
export  gpio3  gpiochip0  gpiochip504  unexport
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio3/direction 
in
```

See the input value.
This is zero since switch is not pushed.
```
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio3/value 
0
```

Push and hold the switch.
The value should be "1" while the switch is pushed.
```
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio3/value 
1
```


When finished, unregister the GPIO pin.
```
pi@raspberrypi:~ $ echo 3 > /sys/class/gpio/unexport
pi@raspberrypi:~ $ ls /sys/class/gpio/
export  gpiochip0  gpiochip504  unexport
```

## Go implementation with periph.io
Take a look at [Read tactile switch state](https://github.com/oklahomer/go-raspi-training/#cross-compile) to cross-compile and transfer the example code.
The corresponding Go code is located at [/example/002-read-tactile-switch/main.go](https://github.com/oklahomer/go-raspi-training/blob/master/example/002-read-tactile-switch/main.go).

Login and execute the transferred executable.
```
pi@raspberrypi:~ $ ./main 
2020/05/09 04:44:38 Start
2020/05/09 06:25:20 GPIO pin level changed from false to true
2020/05/09 06:25:25 GPIO pin level changed from true to false
^C2020/05/09 06:25:30 Stopped
```

While its execution, the input value can directly be checked as below:
```
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio3/value 
0
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio3/value 
1
pi@raspberrypi:~ $ cat /sys/class/gpio/gpio3/value 
0
```
