// gostress -seed 1411172 -want 8
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 4, 9, 1, ); got != 8 {
fmt.Printf("f1_ssa(9, 4, 9, 1, ) = %v, wanted 8\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 int8, a3 uint64, a4 uint, ) int64 {
switch {} // prevent inlining
a3--
v1 := 0 + (uint8(2) << 2) >> (a3 + a3) ^ uint8(1) << (0 ^ 0) | 2 - 3 | 1 - uint8(1) << 2 & 3 // uint8
_ = v1
a2 -= a2 | (a2 ^ a2 + a2 - a2) >> (uint32(0) << v1 * 0 ^ 1) & (a2 + a2) | (a2 & a2) << (a3 >> a3) // int8
a4 = a1 >> (uint32(3 ^ 3) << (a3 | a3) & 2 & 1 | 1 * 3 & 2) // uint
v1 = v1 << (((a1 ^ a1) << v1) >> (2 + 1 - 0) | a1) // uint8
v1 *= (v1 ^ v1 ^ v1 * v1 | v1 * v1 + v1 * (v1 ^ v1) >> (a4 + a4) - v1 - uint8(1)) >> 0 // uint8
a1 = a4 ^ a4 * a4 | (a4 & a4 ^ a4 + a4) >> (a3 - 2 + a3) - (a4 * a1 & (a1 << a3) >> (v1 >> 3)) // uint
a4 -= a1 // uint
v1 = v1 // uint8
return 0 & 2 | 1 * int64(2) >> v1 * 3 - 1 & int64(1) >> v1 & 1 | 3 + int64(1) << (v1 | v1)
}
