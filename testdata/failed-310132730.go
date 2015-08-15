// gostress -seed 310132730 -want 20
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 1, 5, ); got != 20 {
fmt.Printf("f1_ssa(2, 1, 5, ) = %v, wanted 20\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 int, a3 int8, ) int8 {
switch {} // prevent inlining
a3 -= a3 & 2 // int8
v1 := int64(0) << (uint32(3) << (a1 + a1) + 1 & 3) * int64(1) >> (((uint16(2) >> a1) << 0) >> (2 ^ 1 & 1)) | int64(3 + 1 | 0 - 1) << (1 + 0 & 0) * int64(1) >> 0 & int64(1) << (3 ^ 1) // int64
_ = v1
a1 *= a1 * ((a1 | uint64(v1)) << 1) & (a1 ^ a1 >> 1 | a1 - a1) >> (2 | 1 | uint(0) << a1 & 1 * 1) // uint64
v2 := a2 | a2 + (a2 ^ a2) - a2 >> 1 - a2 & a2 & a2 ^ (a2 * a2) + a2 + a2 & a2 * a2 + a2 >> 0 + a2 << 2 - a2 + a2 // int
_ = v2
a3--
v3 := uint((2 ^ 2)) // uint
_ = v3
v3++
v1 = (v1 | v1) >> a1 ^ v1 ^ v1 + v1 >> v3 | (v1 >> (1 & 1)) >> (a1 << 1 * a1) - v1 + (v1 & v1 * v1 ^ v1 >> a1 * v1) << (0 | 1 * 2 ^ uint16(2 - 2) << (1 - 1) | (uint16(1) << v3) + 3 ^ 3 * 1) // int64
v3 -= (((v3 >> a1) | v3 >> (a1 & a1 | a1 * a1)) << ((3 * uint8(v1) << (v3 + v3)) >> (1 ^ (uint8(2) << 3) >> (uint32(0) << 1)))) // uint
return a3 | (((a3 >> v3) << (2 * 0) + a3 << v3) >> (3 >> 1 & 1 * 1 | 0)) << (uint32((2 & 0)) >> (0 + 1) ^ 2)
}
