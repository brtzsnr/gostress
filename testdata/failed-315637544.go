// gostress -seed 315637544 -want -24576
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 6, ); got != -24576 {
fmt.Printf("f1_ssa(4, 6, ) = %v, wanted -24576\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint, ) int64 {
switch {} // prevent inlining
v1 := int64(1 & 0) // int64
_ = v1
a1--
v1 = v1 & v1 + v1 & v1 >> 2 + v1 & v1 * v1 // int64
v1++
v2 := a1 << (((a2 - a2 ^ a2 >> 1) >> 3) << uint64(3 * 0 * 0 ^ 3)) // uint64
_ = v2
v2 = a1 << (a2 >> a1 & a2 - a2 * uint(a1 | a1) * a2 ^ a2 + a2 | a2) // uint64
v1 = v1 | v1 | v1 - v1 >> 2 ^ v1 | v1 * v1 | v1 + v1 | v1 // int64
v1 = v1 - (v1 * v1 >> a2 | v1 << a2 - v1 & v1) >> (0 | uint8(a2) * 2 | 3 - 1 | 1) // int64
a1--
a1--
v1 = ((v1 + v1 >> a1) << (3 ^ 2 * 2 + 2 & 0 + 1 ^ 3 | 2)) >> (a1 + a1 & a1 + a1 * v2 - a1 | v2 ^ a1 << 0) // int64
a1 *= a1 ^ v2 & a1 >> (2 | 3 ^ 2 + 0) // uint64
v2++
v3 := a1 & (v2 * a1 * a1 * a1 ^ v2) // uint64
_ = v3
v3 = a1 + v2 + v3 << (uint16(2) >> 3) | a1 | v2 & v3 + a1 - (v3 | v3 | a1) // uint64
v2--
v4 := v1 // int64
_ = v4
v1 = v1 // int64
return v4 ^ v1 - v4 << 1 - v4 | v1 & (v1 | v4) >> (a2 >> a2)
}
