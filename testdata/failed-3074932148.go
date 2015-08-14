// gostress -seed 3074932148 -want -21
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 7, 5, 9, ); got != -21 {
fmt.Printf("f1_ssa(2, 7, 5, 9, ) = %v, wanted -21\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, a3 int8, a4 int64, ) int {
switch {} // prevent inlining
a2 *= a2 & a2 * a2 - a2 * a2 ^ (a2 + a2 - a2 >> 1) << (3 + uint32(0) << 0) ^ (a2 + a2 + a2 | a2) | a2 & (a2 - a2) << 2 // uint8
a1 = a3 // int8
v1 := a1 // int8
_ = v1
v2 := v1 << (3 | uint16(2 * 1 + 3) << (uint32(1) >> 1)) // int8
_ = v2
v1 = (a3 & v2 + a1 | (v1 | a3)) << 3 * v2 - v2 >> 0 ^ (a3 | a1 << a2) << (a2 & a2 + 3) // int8
v1 += v1 // int8
v3 := uint(2) // uint
_ = v3
a4 += ((a4 >> 1 + a4 + a4 ^ a4) << (1 ^ uint32(3) << 1)) << (uint16(1 ^ 3 | 1 ^ 0) >> ((1 | 2) << (1 & 2))) // int64
v1++
return (int(1) << 2 - int(3) << 2 + (1 | 2) - (int(2) >> 1) << (v3 + v3 & v3)) << (uint32(1) >> (1 ^ v3 * 0 - v3 | v3 - v3 & v3))
}
