package main

import "fmt"

func main() {
 
  fmt.Println("We are going to find out what each number you put in is in terms of Kelvin, Celsius, and Fahrenheit.")
 
  var fahrenheit float64
 
  fmt.Println("enter a number for conversion")
 
  fmt.Scan(&fahrenheit)

  celsius:= (fahrenheit-32)/1.8
  kelvin:= 273.15+celsius

  fmt.Println("That number is",fahrenheit,"degrees Fahrenheit",celsius,"degrees Celsius",kelvin,"degrees Kelvin")
  
}
