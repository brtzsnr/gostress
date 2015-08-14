package main
import "fmt"
func main() {
fmt.Println(f1())
}
func f1() uint {
switch {} // prevent inlining
v1 := uint8(4) // uint8
v2 := uint8(3) // uint8
v3 := uint(2) // uint
return uint(5) << (uint(2) + uint(4) >> (v1) - uint(v1) >> (uint16(0)) >> (uint16(v2)) - v3)
}
