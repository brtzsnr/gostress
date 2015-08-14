// gostress -seed 2877427100 -want 1340
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 1, 0, ); got != 1340 {
fmt.Printf("f1_ssa(2, 1, 0, ) = %v, wanted 1340\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint64, a3 int, ) int {
switch {} // prevent inlining
v1 := uint8(3) // uint8
_ = v1
v1 -= v1 // uint8
a3 -= a3 + a3 | a3 // int
v2 := a1 // int64
_ = v2
a1 *= (a1 + a1) // int64
v2 -= (v2 ^ a1 & 1 * 2) << 1 - a1 << ((v1 + v1 + v1) << (a2 - a2)) // int64
v1 += uint8(3) // uint8
a3 += (a3 | a3) >> (0 * 1) * a3 - a3 | a3 - a3 * 2 & a3 * a3 | a3 >> ((a2 & a2 * a2 | a2) << (2 * 1 & 0 | 0)) // int
a3 += int((v1 << (2 * 2) - v1 * v1 - v1)) // int
a1 -= v2 ^ a1 & a1 + v2 + a1 & a1 | v2 + 2 | v2 << a2 - a1 | (a1 << 2) // int64
return a3 | a3 ^ a3 | a3 * a3 | a3 - a3 | a3 | 3 * a3 >> 3 & (a3 ^ a3 >> a2) << (a2 | a2 | a2 | a2)
}
