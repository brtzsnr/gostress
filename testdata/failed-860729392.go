// gostress -seed 860729392 -want 53
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, 2, ); got != 53 {
fmt.Printf("f1_ssa(5, 2, ) = %v, wanted 53\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, ) int8 {
switch {} // prevent inlining
a1 -= a1 - a1 ^ a1 & int8(3) + a1 | (a1 ^ a1) << (0 | 0) | a1 | a1 + a1 << a2 ^ a1 ^ a1 << 3 + a1 - a1 + a1 + (a1 * a1 * a1 - a1) << (3 - 2 + uint16(2) << 2) * (a1 - a1 & a1) - (a1 >> a2) ^ (a1 * a1) << 2 // int8
v1 := a1 // int8
_ = v1
v1 = v1 + (v1 ^ v1) ^ 1 | v1 | v1 ^ a1 | v1 - v1 - v1 * a1 | (v1 - v1) << (0 << 2) + a1 + v1 * v1 & v1 | v1 * v1 + v1 * v1 | v1 - (v1 - v1) << a2 * (v1 << (2 | 3)) << uint16(0 ^ 0) // int8
v1--
a1--
v1 = v1 // int8
v1--
v1 *= (v1 - v1 & v1 >> 2 & a1 | 2 + v1 - v1 + v1 & v1 & v1 & a1 >> a2 - v1 | a1) ^ (int8(3) & v1 * 1 & v1 & (v1 ^ v1) >> (2 * 1)) >> (0 * 0 & uint16(3) << 1 + 0 * 2) // int8
v1 = a1 | v1 - v1 & (v1 & a1) ^ v1 ^ v1 * a1 << 2 | v1 // int8
return (v1 + a1 | a1) << 3 - v1 - v1 & (v1 >> 2 * v1 << 1 * v1 & v1 ^ v1 - v1) << 1
}
