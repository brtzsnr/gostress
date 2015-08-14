// gostress -seed 1538226765 -want 12
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, 2, 1, 1, ); got != 12 {
fmt.Printf("f1_ssa(9, 1, 1, 9, ) = %v, wanted 12\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 int, a3 int8, a4 int, ) int {
switch {} // prevent inlining
a1 -= (a1 ^ a1 + a1) << (0 & 3 & uint32(2) >> 0) * a1 + a1 << ((1 ^ 1) << 1) ^ a1 // uint
v1 := uint64(1) // uint64
_ = v1
v1 *= (v1 * v1 & v1 >> 0) >> 3 + (v1 & v1) & v1 ^ v1 & uint64(a1) * v1 // uint64
a1 *= (2 | a1 & a1 | a1 | (a1 & a1) >> (v1 | v1) - (a1 + a1) ^ a1) >> ((uint8(0 ^ 3) >> (a1 - a1 * a1 + a1)) << (1 + 0 & 2 - 2 & 0)) // uint
v1 *= (v1 >> (a1 | a1 - a1)) << (a1 & 0 | a1 | (a1 ^ a1)) ^ v1 // uint64
v1++
a2 = (a2 >> v1) + a2 << v1 // int
v1 -= v1 & (0 | v1 - v1) ^ v1 + v1 + v1 ^ v1 // uint64
v1--
v1 = v1 >> (v1 - v1 - v1 * v1 ^ v1 << a1 * (v1 << 0) << (1 * uint16(2) << 1)) // uint64
return a2 + a4 - a4 & a2 * 1 + int(0) - a4 & a4 << 2 | a2 - a4
}
