// gostress -seed 204495578 -want 2
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 3, 3, ); got != 2 {
fmt.Printf("f1_ssa(1, 3, 3, ) = %v, wanted 2\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, a3 int64, ) int {
switch {} // prevent inlining
v1 := uint64(3) // uint64
_ = v1
v1 -= (v1 - (v1 + v1) | (v1 | v1) & v1 & v1 ^ v1 >> 3 * v1) // uint64
v1 *= v1 ^ (v1 | v1) | v1 + v1 * v1 - v1 >> (0 + 3) + v1 * v1 + v1 | v1 | v1 // uint64
a3++
v2 := (v1 ^ v1 * v1 ^ v1) << (uint16(3) << (uint32(2) >> (3 + 3))) ^ v1 >> (a2 ^ a2) ^ v1 | v1 ^ v1 + v1 + (v1 & v1 * 0) >> (uint32(1) << (1 | 1)) - v1 >> (3 * 2 * 0) // uint64
_ = v2
v3 := a1 // int8
_ = v3
v4 := (v1 & v2 & v1 ^ v1 + v2 + v2 - v1 + v2 << ((v2 & v1) >> (0 | 0) + v2 - v2 ^ v2)) // uint64
_ = v4
a1 += v3 + (a1 + a1 << 2 - v3 + a1) >> (a2 - a2 & a2 - a2 + a2) - v3 << ((a2 - a2) >> (a2 ^ a2) * a2 & a2 * a2) // int8
v2--
return (3 ^ int(1) << 2 | int(v2) & int(2) >> 2 + int(0) << a2 + 3) >> a2 + 1
}
