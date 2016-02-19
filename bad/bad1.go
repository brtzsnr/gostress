// gostress -seed 6068
package main
import "fmt"
var b2i = map[bool]int{true:1}
func main() {
fmt.Println(f1_ssa(0, 3, 0, 1, true, 2, 1, 2, ))
}
func f1_ssa(a1 int8, a2 uint8, a3 int8, a4 uint, a5 bool, a6 int8, a7 uint8, a8 int, ) uint8 {
switch {} // prevent inlining
a4--
a8 += int(3 ^ 0) | a8 // int
if a5 {
a1 -= a3 | a6 + ((1 * a3) ^ (a1 * a3) << 1) // int8
a2 -= a2 + a2 // uint8
a2 -= a2 - (a2 + a2) // uint8
v1 := uint64(1) // uint64
_ = v1
a1 = a6 - ((a3 & 1) ^ a1 * a3) // int8
v2 := a1 ^ ((a1 + a6 >> (uint16(3) >> a2)) - (0 | (a1 - a1))) // int8
_ = v2
v1 = v1 // uint64
a4 += (a4 * (a4 * (a4 | a4))) - (a4 >> 0) // uint
}
v3 := a2 + a2 | (a7 >> a7) * (a7 ^ a2 << a4) >> uint64(a8) // uint8
_ = v3
v3 += v3 << a4 // uint8
a1 *= a3 - ((a3 << 2 * (3 | a3)) + ((a6 >> 0) + (a1 & a6))) // int8
a3 *= a1 // int8
v3 += v3 ^ v3 ^ v3 + v3 << ((3 * (3 * 1)) << (1 ^ 3)) // uint8
return v3
}
