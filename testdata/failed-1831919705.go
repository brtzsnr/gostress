// gostress -seed 1831919705 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 7, 9, 9, 1, ); got != 0 {
fmt.Printf("f1_ssa(4, 7, 9, 9, 1, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 int, a3 uint, a4 uint, a5 uint64, ) int64 {
switch {} // prevent inlining
a3 *= uint(a2 * a2) // uint
a5--
a1 -= a1 - a1 * a1 + a1 << 0 & (a1 >> a5) << a5 * a1 - a1 ^ a1 | a1 & a1 & a1 // int64
a3 += a4 + a3 & (a3 << 2 - a3 ^ a4 - a3) >> (3 + 1 | 0 * 3 + 3 ^ 3) // uint
v1 := uint8(3) // uint8
_ = v1
a2 *= a2 >> 2 & a2 & a2 | a2 & (a2 >> a4) | a2 * a2 | a2 & (a2 ^ a2 | a2 - a2) << a5 * a2 >> 3 + a2 | a2 & a2 | a2 & a2 // int
v1 *= v1 // uint8
v2 := a4 // uint
_ = v2
a5 = a5 // uint64
return a1 >> (v1 & v1 - v1) | int64(3) & a1 | a1 & a1 + ((a1 + a1) >> a5) >> (1 - 0) + a1 * a1 ^ a1
}
