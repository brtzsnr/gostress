// gostress -seed 1680219480 -want 2
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 5, 0, ); got != 2 {
fmt.Printf("f1_ssa(2, 5, 0, ) = %v, wanted 2\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint64, a3 uint, ) int64 {
switch {} // prevent inlining
a2 -= (a2 - a2 & a2 ^ a2 ^ a2) >> (uint16(1 - 0 - 0 ^ 1 | 3 - 2) << (3 ^ 0 ^ 1 ^ 0 | 1 * 1) & 2) // uint64
a3 += (a3 ^ (a3 ^ a3) & (a3 >> 3)) >> (2 + 2) & (((a3 << a3) ^ a3 * a3 - a3) >> (2 * uint32(2) << a2 & 3 ^ 2)) << (3 ^ 0 | 2 | 3 ^ 0 | 2 - uint32(int(1) >> (1 | 2))) // uint
a1 = (a1 >> (3 ^ 0 | 1 & 0 ^ 2 ^ 3 - uint32(3) >> 2 - 2 + 2 & 0 & 3 & 1)) // int8
a3 += (a3 ^ a3 * a3 * a3 * a3 * a3 - a3 * a3 - a3 * a3 | a3 - (a3 & a3)) << (a2 ^ a2) // uint
v1 := uint8(3) // uint8
_ = v1
v2 := (a2 | a2 >> a3 + a2 + 0) + (a2 * a2 - a2) << (a2 & a2 | a2) & uint64(a1) + a2 << 0 | a2 * a2 // uint64
_ = v2
a3 = (a3 >> 1 - a3 ^ a3 & a3 + 0 * a3 ^ a3 * a3) << ((a3 | a3) >> 1 - a3 - a3 + uint((v2 | v2 + v2))) // uint
v1++
v2 -= 0 * a2 + a2 ^ v2 + v2 >> 1 | 1 & a2 * a2 + v2 * v2 // uint64
return int64(2) >> a3 & 3 ^ 3 ^ int64(0 + 0) >> (0 & 3) + 0 & 2 * 0 | 1 + int64(2 ^ 3) >> (1 * 2) * 2 & 3 | 3 ^ int64(2 & 2) << a3 ^ int64((1 * 3 ^ 0)) << (v1 << 0 + v1 - v1 - v1) - int64(3 - 2) >> 2 + int64(0 * 0) >> v2 & 3
}
