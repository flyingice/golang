/*
Exercise 2.1:
Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, 
where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1C.
*/
package main

import "fmt"

type Celsius float64
type Kelvin float64

const AbsoluteZeroC Celsius = -273.15

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

func main() {
  fmt.Println(CToK(35.8));
  fmt.Println(KToC(70.2));
}
