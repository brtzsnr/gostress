// gostress -seed 991122110 -want 65536
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(3, 7, ); got != 65536 {
fmt.Printf("f1_ssa(3, 7, ) = %v, wanted 65536\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint8, ) uint64 {
switch {} // prevent inlining
a1 += ((a1 & a1 & 2 * a1 ^ a1 + a1 >> a2) << (uint16(0) << (1 + 1) | 1 * 2 * 3 | 0)) << (a2 | a2 & a2 | a2 >> 3 + a2 + a2 * a2 + a2 - uint8(2) - (a2 - a2) << (a2 | a2)) // uint64
a2 -= a2 & a2 << 2 ^ a2 & a2 & a2 & a2 - a2 // uint8
a1++
a2++
v1 := int64(0 + 1 + 1 - 0) // int64
_ = v1
a1 *= a1 // uint64
v1 -= v1 // int64
v2 := a2 << (1 * 0 * 3 * 1) // uint8
_ = v2
a2 = (a2 + (v2 ^ a2) << (uint16(0) << v2) + uint8(0) | v2 | v2 + v2) >> (v2 & v2 ^ v2 ^ v2 & a2 * v2 - a2 & a2 * (a2 << a1)) // uint8
v2 *= v2 // uint8
return ((a1 - 1) >> (uint16(3) >> 0) ^ (uint64(1) << 1) >> (3 - 1) * (a1 ^ 1 + a1) >> a2) << a1
}
