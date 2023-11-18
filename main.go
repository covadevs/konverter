package main

import "fmt"

const boilingKelvin float32 = 373.0

func main() {

	fmt.Printf("The boiling temperature in Kelvin is %g and in celsius, it is %g", boilingKelvin, kelvinToCelsius())

}

func kelvinToCelsius() float32 {
	return boilingKelvin - 273
}
