// gostress -seed 1126823486 -want -282574488338432
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 2, ); got != -282574488338432 {
fmt.Printf("f1_ssa(4, 2, ) = %v, wanted -282574488338432\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, ) int {
switch {} // prevent inlining
a2--
v1 := 0 // int
_ = v1
v1--
v1 *= (v1 + v1) & v1 | v1 * v1 ^ 1 & v1 + (v1 >> (a2 | a2) + (v1 * v1) << a2) >> (3 ^ 2 & 0 - 1) ^ v1 // int
v1 -= ((3 + v1) >> (a2 << 1)) >> (3 ^ 1) ^ (v1 & 2 + v1 >> 0) << (uint32(0 ^ 3) << 3) + v1 - v1 << (0 ^ uint16(uint(2) << 3)) // int
a1 -= a1 - (a1 & a1) >> (1 - 1) - a1 & a1 | a1 * a1 & a1 + a1 ^ (a1 ^ a1) >> 0 + a1 - 2 ^ a1 ^ a1 * a1 * a1 | 2 | a1 | a1 & a1 | a1 | a1 ^ a1 | a1 ^ 1 * a1 // int8
a2++
v1 *= v1 ^ (v1 ^ v1 - v1 & v1) << 2 - (v1 << a2) >> 0 + v1 | v1 & v1 | v1 ^ v1 // int
a2++
a2 -= a2 + a2 * a2 ^ a2 | 2 * a2 * (a2 >> (uint16(1) << 2) ^ 1 + a2) & a2 * a2 & a2 + a2 + a2 * a2 | (a2 ^ a2) * a2 ^ a2 + a2 // uint8
return v1
}
