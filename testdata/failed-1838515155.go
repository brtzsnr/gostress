// gostress -seed 1838515155 -want 34
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 9, ); got != 34 {
fmt.Printf("f1_ssa(2, 9, ) = %v, wanted 34\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint, ) int8 {
switch {} // prevent inlining
a1 = ((a1 & a1 ^ a1) << (3 + 3)) << 3 - (a1 - a1 & a1) * a1 >> 3 & a1 + a1 - a1 + a1 - a1 // int64
a2 *= a2 >> a2 // uint
a1 *= a1 << (uint64(1 & 2) >> ((a2 & a2 - a2 | a2) << (0 * 0 ^ 1))) // int64
v1 := uint8(1) // uint8
_ = v1
v2 := ((a1 - a1) << (0 + 1 * 0 + 3) ^ (a1 << 2 + a1 & a1) >> (v1 * v1 | v1 << 3) * a1 | a1 << v1 * a1 >> a2 + a1 + a1 << 3 + a1) << v1 // int64
_ = v2
v3 := v2 // int64
_ = v3
a2 *= a2 // uint
a2 *= a2 - (a2 << 2) >> (uint16(3) << 0) & (a2 + a2 * a2) - (a2 | a2 & uint(3) << 0 ^ a2 - a2 * a2 * a2) + (a2 >> 2 | a2 ^ a2 ^ a2) << ((uint16(0) << 1) ^ 3 & uint16(0) >> (1 & uint32(3) << 1)) // uint
a1 += a1 ^ (v3 << (0 + 2)) << 0 - v2 | v2 // int64
v1++
return 0 & 3 + (int8(3) >> 2 + 3 ^ 0 ^ int8(1) >> 3 + 1 * 2) * 0 & 3 & 2 & int8(3 & 2) >> a2 + (3 - 1 * 3) | 3 & int8(2) << (a2 & a2) * int8(2) << 3 + int8(2) >> 0
}
