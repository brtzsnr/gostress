// gostress -seed 2452124262 -want 263424
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(5, 3, 6, ); got != 263424 {
fmt.Printf("f1_ssa(5, 3, 6, ) = %v, wanted 263424\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint, a3 int8, ) int {
switch {} // prevent inlining
v1 := ((2 * 0) * 2 & int(2) << 3 - 3 | 2 - 0) << 1 // int
_ = v1
v2 := ((uint64(3) >> 2 ^ 0 | 3) >> 2) >> (3 & 0 * 1) * uint64(0) & 2 ^ 2 - uint64(0) << 3 * 2 + (2 & 1 ^ uint64(2) >> 0) * 1 ^ 3 & 3 | 3 & uint64(2) << 0 & 1 + uint64(3) >> (a2 + a1) - uint64(2 & 3) << (2 ^ 3) - 2 + 0 | 0 ^ 2 - uint64(1) << 3 | 1 + 3 // uint64
_ = v2
a1 -= ((a2 ^ a2 ^ a1) << (uint32(0 ^ 1 ^ 3 & 2) >> (2 | 2) + 3 + 3 & 0 & 0 & 0 * uint32(3) << 0)) // uint
a2++
a3 *= (a3 | a3 ^ 3 ^ a3 * a3 * a3 - a3) >> (uint32(1 ^ 0) >> (1 - 1) * 2 ^ 0 ^ 3 ^ 1) - a3 << (3 - 3 ^ 0 + uint16(1) >> 1) + (a3 | a3 + a3 << 2 + a3 << 2 - (a3 * a3 ^ a3)) // int8
v2 -= (v2 | v2) // uint64
v3 := ((v1 & 3 + v1 >> v2 * v1) << (uint8(2 & 3) << (a2 | a1) + 3 - 3 | 0) ^ v1 >> (a2 - a2 - a2 >> 1 | a1 - a1 & a2)) << 3 // int
_ = v3
v2 *= v2 ^ v2 << a1 * v2 + v2 - v2 ^ v2 & v2 * (v2 - v2) >> (3 + 0) ^ v2 * v2 & v2 ^ v2 ^ v2 ^ v2 * (v2 << v2) >> 2 - v2 * v2 >> (2 * 2 - 2) + (v2 + v2) >> (uint16(3) >> 3) | v2 * 0 * v2 << 1 + v2 << (v2 >> a1 & v2) // uint64
return (v1 & v3 << a2 & v1 * v3 & v1 * ((v3 * 3 * v1 << 1) >> (1 | 3))) >> a2
}
