// gostress -seed 1959132127 -want 51
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(3, 4, 4, 8, ); got != 51 {
fmt.Printf("f1_ssa(3, 4, 4, 8, ) = %v, wanted 51\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint, a3 int, a4 uint64, ) int64 {
switch {} // prevent inlining
a3 = a3 | a3 + a3 // int
a4++
a4 -= (a4 << 1) >> (3 & 2 - 1) + a4 - a4 * a4 * a4 * a4 + a4 // uint64
a4--
a3++
a4--
v1 := a3 * a3 // int
_ = v1
v1 *= v1 // int
v1 *= v1 | v1 // int
v2 := ((v1 | v1 * a3 + v1) << (1 * 1 * 1 * 1)) & v1 - int(0) << (1 | 2 * 2) // int
_ = v2
return (a1 ^ a1) | (a1 << a2) >> (uint32(2) >> 3) ^ a1 & a1 | (a1 ^ 1) ^ (a1 + a1) >> (0 + 1) | (a1 << a2 ^ a1) >> a2
}
