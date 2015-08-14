// gostress -seed 7002275 -want 5
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 4, 9, ); got != 5 {
fmt.Printf("f1_ssa(4, 4, 9, ) = %v, wanted 5\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint64, a3 int, ) int64 {
switch {} // prevent inlining
a2 *= (a2 | a2) & a2 << a1 - a2 >> 0 + a2 << (2 ^ 0 - 1) ^ a2 >> (a1 * 1 | a1 * a1 & (a1 << 0) >> (1 | 1)) // uint64
a2 += a2 >> (0 + 1 ^ 2 * 2 | 0) & a2 ^ a2 | a2 * a2 + a2 ^ a2 // uint64
a3--
a2 += a2 * a2 ^ a2 | a2 * a2 & a2 // uint64
a2 -= uint64((a1 << 2) >> (1 | 0) & a1) + a2 & a2 - a2 & a2 + (a2 << 1) << (0 << 0 & 2 * 1) // uint64
v1 := int64(1) // int64
_ = v1
v2 := a1 // uint8
_ = v2
v3 := a2 << 3 + a2 | a2 & a2 >> 3 - a2 * a2 * a2 - a2 | a2 + 2 * a2 | a2 - a2 * a2 - a2 * a2 * a2 << 1 // uint64
_ = v3
v1 *= v1 // int64
return 0 ^ v1 + v1 ^ v1 + v1 + v1 - v1 ^ v1 * ((v1 & v1) << a1 + int64(uint(3) >> 1)) >> v2
}
