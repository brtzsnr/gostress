// gostress -seed 2248311269 -want 2
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 7, 1, ); got != 2 {
fmt.Printf("f1_ssa(6, 7, 1, ) = %v, wanted 2\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint, a3 int8, ) int64 {
switch {} // prevent inlining
a1 = (a1 * a1 >> 1) >> (a1 - a1 >> 2 | a1 | a1) * (0 | a1) >> (uint32(2) << 1) * uint64(3 ^ 3) ^ (a1 | a1) ^ a1 | a1 & a1 | a1 // uint64
a1 -= (a1 - a1 >> a1) >> ((1 << ((uint32(0) << 3) >> 2)) << a1) // uint64
a3 *= (a3 - a3 | a3 + a3 | a3) * a3 + a3 ^ a3 & a3 + a3 | a3 - a3 >> (a1 << 1 | a1) // int8
v1 := int64(3) // int64
_ = v1
a1--
v1 += v1 >> ((0 << 1) << (a1 - a1) | 3) // int64
a2--
v1 = (v1 ^ int64(a2) | v1) << (2 + 2 | uint32(3) << (uint8(2) >> 3) - 3 + 1 * uint32(a2) * uint32(2) << (a1 * a1)) // int64
a2 = (2 + a2) >> uint8(a1) & a2 - a2 * 0 + a2 ^ a2 >> 3 & a2 >> (2 | 1) + a2 >> 3 - a2 // uint
return v1 ^ v1 << 1 | v1 >> a2 * v1 * v1 * 2 & v1 + v1 ^ v1 + (2 + v1 * v1 & v1)
}
