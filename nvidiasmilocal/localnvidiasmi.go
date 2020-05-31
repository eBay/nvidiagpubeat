package main

import "fmt"

func main() {
	fmt.Println("name, driver_version, count, index, fan.speed [%], memory.total [MiB], memory.used [MiB], utilization.gpu [%], utilization.memory [%], temperature.gpu, power.draw [W], power.limit [W], clocks.current.graphics [MHz], clocks.current.sm [MHz], clocks.current.memory [MHz], pstate")
	fmt.Println("Tesla P100-PCIE-16GB, 418.87.01, 4, 0, [Not Supported], 16280 MiB, 1628 MiB, 10 %, 10 %, 25, 24.80 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
	fmt.Println("Tesla P100-PCIE-16GB, 418.87.01, 4, 1, [Not Supported], 16280 MiB, 3256 MiB, 30 %, 20 %, 25, 25.05 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
	fmt.Println("Tesla P100-PCIE-16GB, 418.87.01, 4, 2, [Not Supported], 16280 MiB, 1628 MiB, 20 %, 10 %, 24, 26.26 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
	fmt.Println("Tesla P100-PCIE-16GB, 418.87.01, 4, 3, [Not Supported], 16280 MiB, 3256 MiB, 70 %, 20 %, 24, 25.28 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
}
