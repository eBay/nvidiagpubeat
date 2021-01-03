package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		utilization()
	} else {
		arg := os.Args[1]
		if arg == "utilization" {
			utilization()
		} else {
			process()
		}
	}

}

func utilization() {
	fmt.Println("name, driver_version, count, index, fan.speed [%], memory.total [MiB], memory.used [MiB], utilization.gpu [%], utilization.memory [%], temperature.gpu, power.draw [W], power.limit [W], clocks.current.graphics [MHz], clocks.current.sm [MHz], clocks.current.memory [MHz], pstate")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:08:00.0, 0320218165889, GPU-78f90e78-39a0-4f40-fcbc-0adf3598c166, 418.87.01, 4, 0, [Not Supported], 16280 MiB, 1628 MiB, 10 %, 10 %, 25, 24.80 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:0B:00.0, 0320218176885, GPU-d1229c61-babc-aebe-ff8f-6dc94386640c, 418.87.01, 4, 1, [Not Supported], 16280 MiB, 3256 MiB, 30 %, 20 %, 25, 25.05 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:16:00.0, 0320218166179, GPU-eb5e8723-4a49-98f5-8e77-21b06537da8a, 418.87.01, 4, 2, [Not Supported], 16280 MiB, 1628 MiB, 20 %, 10 %, 24, 26.26 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:19:00.0, 0320218176911, GPU-b884db58-6340-7969-a79f-b937f3583884, 418.87.01, 4, 3, [Not Supported], 16280 MiB, 3256 MiB, 50 %, 50 %, 24, 25.28 W, 250.00 W, 405 MHz, 405 MHz, 715 MHz, P0")
}

func process() {
	fmt.Println("gpu_name, gpu_bus_id, gpu_serial, gpu_uuid, pid, process_name, used_gpu_memory [MiB], used_gpu_memory [MiB]")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:16:00.0, 0320218176947, GPU-bb7f65ee-acdb-7efd-0f32-73699400b86e, 240930, python, 10 MiB, 15 MiB")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:0B:00.0, 0320218176885, GPU-d1229c61-babc-aebe-ff8f-6dc94386640c, 65808, python, 10 MiB, 15 MiB")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:16:00.0, 0320218166179, GPU-eb5e8723-4a49-98f5-8e77-21b06537da8a, 267414, python, 10 MiB, 15 MiB")
	fmt.Println("Tesla P100-PCIE-16GB, 00000000:19:00.0, 0320218176911, GPU-b884db58-6340-7969-a79f-b937f3583884, 222414, [Not Found], 10 MiB, 15 MiB")
}
