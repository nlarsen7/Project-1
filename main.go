//Nicholas Larsen
//February 13, 2020
//This converts a number to Fahrenheit, Celsius, and Kelvin.
package main

import "fmt"

func main() {
 
  fmt.Println("We are going to find out what each number you put in is in terms of Kelvin, Celsius, and Fahrenheit.")
  //Explains the function
 
  var fahrenheit float64
 
  fmt.Println("enter a number for conversion")
 
  fmt.Scan(&fahrenheit)
  //Scans for Fahrenheit since there isn't a conversion for it

  celsius:= (fahrenheit-32)/1.8
  kelvin:= 273.15+celsius
  //The conversion equations for the others

  fmt.Println("That number is",fahrenheit,"degrees Fahrenheit",celsius,"degrees Celsius",kelvin,"degrees Kelvin")
  //Prints out all three tempuratures for the user
  
}
