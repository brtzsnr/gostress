// gostress -seed 1723120894 -want -2662
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 3, 2, ); got != -2662 {
fmt.Printf("f1_ssa(2, 3, 2, ) = %v, wanted -2662\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int, a3 int, ) int {
switch {} // prevent inlining
a1--
v1 := (int64(int8(3) << 2 & 1) - 3 - int64(0) << 1 ^ 1 ^ 3 ^ 3 + 2) >> (3 + uint16(1) << ((0 & 1) << 1) ^ 0 ^ (uint16(0) >> a1) >> 2) // int64
_ = v1
v1++
a1++
v1 *= v1 - v1 * (v1 & v1 | v1) * v1 * v1 << 2 - v1 - v1 & v1 // int64
a3 += a3 // int
a3 *= (a2 & a2 ^ a2 - a3 | a3 & a2 + a2) ^ (a2 - a2 & a3) << 3 + int(3) // int
v2 := (1 ^ 0) + (int8(3) >> 3) >> (uint(1) >> 3) + 1 & 3 & 3 & (int8(3 ^ 3) >> (3 - 2) | 0 | 1) << 0 // int8
_ = v2
v2--
return a3 * ((a3 | a3) << (1 ^ 3) * a3 << 2 ^ a2 & ((a2 + a3) << (a1 << 0))) >> (uint64(2 ^ 0 | 3) << (a1 << 0))
}
