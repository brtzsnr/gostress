// gostress -seed 3088432620 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 0, 9, 9, ); got != 0 {
fmt.Printf("f1_ssa(8, 0, 9, 9, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 int, a3 uint8, a4 int64, ) uint64 {
switch {} // prevent inlining
a2 -= a2 & a2 >> (a3 & 2) * a2 >> (uint(0) >> 0 + 3 - 1) // int
a3 -= (((a3 >> 1) >> (0 | 0)) ^ (a3 | a3 ^ a3) >> (0 ^ 0 * uint32(3) << 1)) << ((a3 & a3) >> (1 + 3) | uint8(0) & a3 >> 2 - (a3 - a3) >> 1 + a3 ^ 1 + a3 * a3) // uint8
a4 *= (a4 >> 2 * a1 << 3 | a1) << (0 + a3) // int64
a1 += a4 // int64
v1 := ((a4 & a1 + a1 - a4) | a1 - a4 >> a3 * a1) >> (uint64(2 * 1) << (uint16(2) << 2) & uint64(3) + uint64(1) | 1) // int64
_ = v1
v1 -= (a1 - v1 ^ v1 * v1 | 1 + v1 ^ a4 + 1) >> 3 // int64
v2 := int64(a3 & a3 - a3 + a3 & a3 << (2 & 1) | a3) // int64
_ = v2
a2--
v2++
v3 := ((uint64(3) << a3 - 3 * 1 * 3 + 3 ^ 1 & 0) >> (uint(1) << 2)) >> (a3 & a3 << (uint32(0) << 3)) // uint64
_ = v3
return (uint64(v2 & v1) * v3 & v3 | v3 - v3 - v3 | v3 ^ v3 << v3 - v3 & v3)
}
