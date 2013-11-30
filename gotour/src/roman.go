// Example not from Tour but from
// http://stackoverflow.com/questions/1720201/go-examples-and-idioms 
package main 

import (
	"fmt"
	"strings"
)

var numText = "zero one two three four five six seven eight nine ten"
var numRoman = "- I II III IV V VI VII IX X"
var aText = strings.Split(numText, " ")
var aRoman = strings.Split(numRoman, " ")

type TextNumber int
type RomanNumber int

func (n TextNumber) String() string {
  return aText[n]
}
func (n RomanNumber) String() string {
  return aRoman[n]
}
func main() {
  // var i = 5
  for i := 0; i < 10; i++ {
  	fmt.Println("Number: ", i, TextNumber(i), RomanNumber(i))
  }
}
