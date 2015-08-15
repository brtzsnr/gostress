// gostress -seed 81332618 -want 812
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 3, ); got != 812 {
fmt.Printf("f1_ssa(9, 3, ) = %v, wanted 812\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int, ) int {
switch {} // prevent inlining
a2++
v1 := uint8(2) // uint8
_ = v1
v1 -= ((3 * (v1 >> v1) - v1) << 0) >> (v1 - v1) // uint8
a2 += (((a2 - a1 & a1) >> (uint64(1) >> 0 ^ 0)) >> v1 | 3 | a2 & a1 * a1 * a2 + a1 & a2) // int
a2 = ((a2 & a1) << (uint32(2) >> 3)) >> (uint64(0) << ((uint(0) >> 1) << (1 * 3))) + int(uint64(2) << 1) * a1 - a2 | a2 * (a2 | a2) << (0 * 2) | a2 * a1 - a2 | (a2 - a1 - int(0) ^ a2 ^ a1 ^ a1) >> ((v1 >> 0 + v1 * v1 >> 1 | 0 - v1) << (2 * 3 | 3 + uint(3 * 2) >> 0)) // int
a1 -= (a1 - a2 * a2 << v1) >> (uint64(1) >> 2 & 3) & (a2 >> 3) + (a1 + a2) << v1 + a1 * a2 << 2 // int
v2 := a1 // int
_ = v2
v3 := (v2 - (a1 | v2) >> (0 | 3) * a1 ^ a2 | a2 | a2) << (2 + 2 - 1 + 1 - 3 | 1 * 0 + 0 & 2 + 1) - v2 >> 0 | a2 + a2 ^ (a1 << 1) >> (0 << v1) * a1 // int
_ = v3
a2 -= a2 >> (1 + 3) // int
return (v2 ^ v2 & ((a2 + v3) << v1))
}
