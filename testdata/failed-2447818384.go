// gostress -seed 2447818384 -want -7
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 3, 0, ); got != -7 {
fmt.Printf("f1_ssa(2, 3, 0, ) = %v, wanted -7\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint64, a3 int64, ) int {
switch {} // prevent inlining
v1 := a1 << 1 & a1 * a1 & int(3 & a2) + a1 & (a1 - 3) >> uint32(2) - a1 & a1 ^ a1 | a1 - a1 >> (2 | 0 | 1 ^ 2) // int
_ = v1
v1 = v1 >> ((a2 & a2) + (a2 * a2) >> 3) + 0 & (v1 >> 3) - v1 + v1 & v1 | a1 // int
v2 := (uint((2 | 3)) >> 0) ^ uint(1) >> 2 * 3 * 3 ^ 2 | 1 ^ 3 - 3 & 3 & 0 + 0 // uint
_ = v2
v1 *= v1 // int
v3 := a3 // int64
_ = v3
v2++
a3 = a3 // int64
v3 = v3 & v3 + v3 // int64
a2--
v1 -= v1 >> a2 ^ v1 >> v2 ^ v1 ^ a1 ^ v1 ^ a1 * a1 - v1 - 1 + v1 << (a2 & a2) // int
return v1
}
