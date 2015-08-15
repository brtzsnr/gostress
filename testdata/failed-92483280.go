// gostress -seed 92483280 -want -556
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 0, 7, ); got != -556 {
fmt.Printf("f1_ssa(1, 0, 7, ) = %v, wanted -556\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint, a3 int64, ) int64 {
switch {} // prevent inlining
v1 := uint64(3) // uint64
_ = v1
v1 *= uint64(a3) // uint64
v1 -= v1 // uint64
v1 *= v1 // uint64
v1 = ((v1 & uint64(2)) << 3) >> a2 + ((v1 ^ v1) >> (a2 - a2)) >> (2 & (1 >> 2) >> (3 - v1)) // uint64
a2 += a2 & a2 // uint
v2 := a3 // int64
_ = v2
v2 -= (v2 - v2 + a3 + v2 - v2) << (3 * 1 | 1 * 3 * 0 | 1 * 3) + v2 + a3 ^ v2 >> a1 + a3 >> 2 // int64
v2 = v2 - v2 ^ a3 << a1 & (a3 | a3) | v2 >> (v1 | v1 & v1) | v2 | a3 - v2 * a3 & a3 | v2 * v2 | a3 >> 2 // int64
return a3 << a2 - a3 | v2 - v2 >> 0 & a3 ^ v2 + v2 << v1 + v2 * a3 + a3 + v2 + a3 >> a1
}
