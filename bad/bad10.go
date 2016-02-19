// gostress -seed 29823
package main
import "fmt"
func b2i(b bool) int {
if b {
return 1
}
return 0
}
func main() {
fmt.Println(f1_ssa(1, 1, 3, 3, 1, 2, 0, 3, ))
}
func f1_ssa(a1 int8, a2 uint, a3 uint, a4 uint8, a5 int, a6 int8, a7 uint64, a8 int8, ) uint8 {
switch {} // prevent inlining
a6 = a1 // int8
a7 -= a7 // uint64
a5--
a7--
if 0 != (a4 | (a4 + a4) << 1) {
a8++
a7 -= (a7 | (a7 - a7) >> 2) + ((a7 ^ a7 >> (a4 & a4)) + (a7 - a7 & (a7 | a7))) // uint64
a7 -= (a7 & (a7 - a7)) ^ a7 // uint64
a4++
a7--
a8 -= a6 // int8
}
a2 = a2 << (a4 | a4 - (a4 * a4) + ((a4 >> 2) | a4)) // uint
a6--
a2--
return a4 >> 0
}
