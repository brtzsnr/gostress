// gostress -seed 149426415 -want 35
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 5, 1, ); got != 35 {
fmt.Printf("f1_ssa(2, 5, 1, ) = %v, wanted 35\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int8, a3 uint64, ) int8 {
switch {} // prevent inlining
v1 := a3 // uint64
_ = v1
v2 := (a1 + a1 >> 3 | 0 * a1) << uint16(3) | (int(2 & 0) & a1 - a1 + a1 ^ a1) >> (0 ^ 1 * 3 & 0 ^ 0 | 0 & 0 & 1) // int
_ = v2
v2--
v2--
a2 *= (a2 & a2 >> v1 ^ a2 << v1) - a2 // int8
v1 = a3 ^ v1 & a3 & a3 & v1 | a3 + (a3 >> 0) >> (2 * 0) - v1 * a3 // uint64
v1 += (a3 & a3) << (v1 | a3) - a3 - a3 + v1 & v1 * a3 * (v1 - v1) | v1 << (uint32(0) >> 0) - v1 << 0 + v1 + v1 * a3 * v1 * a3 >> 0 // uint64
v2 += (a1 >> uint8(a1) | 1 ^ 3 | a1 ^ a1) << (uint16(2 & 0) >> (v1 * a3) ^ 3 * uint16(a1)) & ((a1 & 2) >> 1) ^ a1 & a1 ^ v2 << a3 + a1 ^ v2 + a1 // int
return int8((int64(0 | 2) >> 1) << 0) - (a2 * a2 - a2 >> 1) >> (0 | 1 - 0) + (a2 + a2 - a2)
}
