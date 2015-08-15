// gostress -seed 1724229112 -want -3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, ); got != -3 {
fmt.Printf("f1_ssa(1, ) = %v, wanted -3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, ) int8 {
switch {} // prevent inlining
a1--
a1 = a1 & a1 & a1 | 0 - a1 + a1 * a1 << (0 | 3 ^ 0 & 1 & 3 + 3) // uint
a1 -= (uint((int(2) << 3 * 1) >> 1) - ((a1 * a1) + a1 ^ a1 - a1) << 1) >> (0 + uint8(1 & 1 ^ 3 * 1 + 2) << ((a1 - a1) >> 0 & (a1 >> 2) << (uint8(3) >> 3))) // uint
v1 := ((uint8(3) >> (a1 + a1 ^ 1)) & uint8(1) >> a1 & uint8(3) ^ uint8(2) | uint8(0) << a1 * uint8(3 - 1 + 2 - 0) >> 1) // uint8
_ = v1
v2 := ((v1 | v1 * v1 | v1) << 2 | (v1 & v1 - v1 << 2) << (3 + 1 ^ 1 & 3) - v1 - v1 + v1 | 3 - v1 * v1 & v1 + v1) << ((v1 - v1 | v1 >> (1 ^ 3)) << 3) // uint8
_ = v2
v1 -= v2 // uint8
v2 += (v1 << (a1 - a1 + a1 ^ a1 ^ a1 >> a1 - a1 & a1 + a1 + a1 & a1)) << v1 // uint8
v2++
v1++
return int8(0) >> a1 + 0 | 1 ^ 0 * 3 - int8(1) << a1 - 1 | 3 + int8(2) >> 2 * 1 * 0 + (int8(2) >> 2) << (1 - 0) - 2 * 1 + 3 & 1 & 3 * 1 * 2 & 3 & int8(0) << 0 - (2 - 1) & ((int8(1) << 2) << (a1 ^ a1))
}
