// gostress -seed 2535428578 -want 30
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, 3, 4, 0, ); got != 30 {
fmt.Printf("f1_ssa(0, 3, 4, 0, ) = %v, wanted 30\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint64, a3 int8, a4 uint8, ) int64 {
switch {} // prevent inlining
a2 = (a2 * a2 | a2) ^ a2 << (2 * 0) & a2 ^ a2 << 3 & a2 // uint64
a2++
a2 -= a2 << (a4 & a4 & a4) & a2 - uint64(1) - a2 ^ a2 - a2 & a2 ^ 3 + (a2 >> 1) >> 0 + a2 - a2 & a2 // uint64
a4 -= (a4 << 2 & a4 * a4 * (a4 + a4)) >> (a4 - a4 + 0 * a4 ^ (a4 << 1) << (a4 - a4)) & a4 // uint8
a1 -= a3 // int8
a4++
a3 *= a1 - (a3 ^ a1) | a1 + a1 - a3 | (a1 * a1 + a1 | a1) // int8
v1 := uint(2) // uint
_ = v1
v1 = v1 // uint
return (int64(3) << 2 | 3 - 1 - 0 | int64(3 & 3 + 2 - 1) << (a4 | 3 ^ a4 & 1)) >> (((uint32(0) << 0) >> (1 - 0) + 0 & 1 & 2 | 2) >> (a2 * 2 ^ a2 >> v1 + (a2 * a2) << (uint16(1) << a2)))
}
