// gostress -seed 255038385 -want 11
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 2, 4, 0, 4, ); got != 11 {
fmt.Printf("f1_ssa(6, 2, 4, 0, 4, ) = %v, wanted 11\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint64, a3 uint, a4 int64, a5 int8, ) int {
switch {} // prevent inlining
a3--
a1 = (a1 & a5 * a5 & a5 - a1 & a1) << a3 // int8
v1 := a4 >> (uint8(3) << 2) | a4 << (1 - 0) + a4 & a4 - a4 | a4 + a4 + 1 + a4 // int64
_ = v1
v1 -= ((v1 - int64(a2 << 0)) << (a2 | a2 ^ a2 | a2 << 1)) // int64
v2 := v1 // int64
_ = v2
v2 += v2 - v2 * a4 & v1 >> (0 << 3) + ((v1 & v2) >> 0) >> (0 ^ 3 + 1 << 0) + v2 & a4 & a4 - v1 + v2 + v2 // int64
a4--
a5 = int8((a3 + a3 & a3) << a2) * a5 // int8
v2 -= a4 >> 0 // int64
return ((1 - int(1) >> a2) << (0 | 0)) << (a3 >> 1 & a3 | a3 & a3 << 0 | a3) ^ 3
}
