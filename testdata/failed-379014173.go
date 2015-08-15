// gostress -seed 379014173 -want 480
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 0, ); got != 480 {
fmt.Printf("f1_ssa(7, 0, ) = %v, wanted 480\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint8, ) int {
switch {} // prevent inlining
a2 = a2 + a2 >> 3 ^ a2 + (a2 + a2) >> 1 | a2 + a2 & a2 - ((a2 >> 1 * a2) >> (uint(1 ^ 2) >> (2 + 3))) >> 2 ^ (a2 >> 3 & a2 * a2) >> (3 + 2 & 2) - a2 & a2 ^ a2 & a2 | (a2 & a2) & (a2 * a2) & a2 * a2 - a2 // uint8
v1 := (int8(3 + 0 & 0) << (uint32(1) << 1) + 1 & 0 + 0 + 2 * 2) << (2 * 2 + 3 | 3 ^ (uint16(3) << 3 & 0 * 3) >> ((a2 >> 1) >> (1 | 0))) * (int8(1) << 0) << 3 + 0 | int8(3) << 0 & int8((3 * 0)) >> 3 ^ (0 + 1) // int8
_ = v1
a1 = (a1 << 0 + a1 & a1 + a1 >> 0 * a1 - a1 | (a1 << 2) >> (1 | 1) ^ a1 | a1 << 1) << 3 ^ a1 & (a1 >> 1) >> a2 * a1 | 3 * a1 - a1 * (2 ^ a1) + a1 ^ a1 - a1 << a2 + (a1 >> 1 | a1) >> (3 + 0 - 0) // int
v1--
v1 *= v1 // int8
v1 *= ((v1 + v1 + v1 * v1 & v1) >> (2 + 2) - v1 - v1 * v1 - v1 & (v1 | v1 & v1) << (a2 * a2)) << (uint32((1 & 2) + 2 + 2 * 2) >> 0 + 0) // int8
v1 -= v1 // int8
v1 += v1 // int8
a1++
v1 += (v1 | (v1 ^ v1)) << ((1 + 0 | 2 * (0 ^ 3) - 3) << 1) // int8
return a1 - a1 | a1
}
