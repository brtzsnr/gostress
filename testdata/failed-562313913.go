// gostress -seed 562313913 -want 3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, ); got != 3 {
fmt.Printf("f1_ssa(7, ) = %v, wanted 3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, ) int64 {
switch {} // prevent inlining
a1 *= (a1 << (1 & 0 * 3 | 2) | a1 - a1 | a1 ^ a1 * a1) << (0 * 1 | 3 | 0 + 0) // uint8
a1--
a1++
a1 = uint8(0) // uint8
a1 *= ((a1 * a1 - a1 << 1) << (3 & 0 * 0)) + a1 << (uint32(2) << 0) - a1 ^ a1 * a1 - a1 >> 2 + a1 ^ a1 ^ a1 >> 2 // uint8
v1 := 1 * 1 // int
_ = v1
a1 -= a1 << (uint16(1) >> a1 | 0 - 2 | 3 * 1 ^ 1 & 2 | 0) // uint8
v2 := uint((uint8(v1) * a1 ^ a1) << (1 + 3 | uint32(2) >> a1) - a1 | (a1 << 0)) // uint
_ = v2
v3 := a1 // uint8
_ = v3
return (int64(3) >> v2) << v2 * 1
}
