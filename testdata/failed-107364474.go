// gostress -seed 107364474 -want -22
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 9, 8, ); got != -22 {
fmt.Printf("f1_ssa(2, 9, 8, ) = %v, wanted -22\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint64, a3 int8, ) int64 {
switch {} // prevent inlining
a2 = (a2 + a2 + a2 * a2 + a2) - ((a2 * a2) >> (1 + 1) + a2 >> (0 | 3)) << (1 & 1 * 3 & 1) ^ a2 + a2 >> 1 - a2 & a2 * (a2 - a2) & a2 - a2 & a2 ^ a2 * a2 | (a2 * a2 & uint64(3)) // uint64
v1 := ((a1 ^ a1 ^ a1 + a1) >> (1 >> (a2 >> 2 * a2)) - a1) // int64
_ = v1
v1 = (((v1 | a1) + v1 & v1 | v1) >> 0) + (v1 << a2) >> a2 & v1 & a1 + a1 - v1 & v1 + v1 - v1 - v1 >> 3 // int64
v2 := (0 * 0 & 2 & 1 | 2 & 1 & 1) & uint(2) >> 0 ^ 0 | uint(a3) | (uint(1) << (uint32(1) >> a2)) << 1 ^ (uint(1 & 1) >> 0) - uint((1 | 1)) >> 0 * uint(2) >> 2 + 3 & 0 + 2 // uint
_ = v2
a2 -= a2 - (a2 * a2 * a2 - a2) << 2 | (a2 & a2 & a2 * a2 + a2) // uint64
a3 *= a3 | (a3 >> 1 | (a3 >> 0) >> (0 + 2)) << 1 - a3 // int8
v3 := (int(0 | 1 & 1) << (v2 << 3 ^ v2 << v2)) << ((a2 >> a2) << (v2 + v2)) - int(3) << (uint8(0) << (2 - 1)) | (1 & 1) & 3 | 0 & 1 & 1 & 0 * 0 ^ int(3) >> 1 | (int(1) >> 1) | 0 - 0 | 2 & 3 & int(2 * 3) >> (3 + 3) * int(2 & 2) >> (a2 ^ a2) * 0 - int(2) >> 0 // int
_ = v3
v3--
a3++
v4 := a1 >> (3 * 2) * v1 * v1 & v1 - a1 | a1 & a1 - a1 ^ v1 - v1 + a1 * a1 - (a1 ^ v1) << (v2 << 1) * v1 & (a1 + a1 + v1 - a1 ^ v1 - a1 ^ a1) | a1 // int64
_ = v4
return ((a1 - v4 & v4 * v1) >> (2 | 2 - 0) & (v4 >> (v2 - v2))) >> (((uint16(0) >> a2) >> (0 & 2)) << (uint32(a3) | uint32(2) << 0) * 2 + 3) | a1 & (a1 >> v2) << 1 - a1 >> (uint8(2) << 0) | v4 + a1 & 1 | v1 ^ a1 * v1
}
