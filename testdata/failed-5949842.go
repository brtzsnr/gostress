// gostress -seed 5949842 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, 2, 7, 3, 0, ); got != 0 {
fmt.Printf("f1_ssa(0, 2, 7, 3, 0, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint, a3 int, a4 uint, a5 int8, ) int {
switch {} // prevent inlining
a1 = int8(1) & a5 // int8
a2 *= a2 // uint
a5++
a2--
a3 = a3 // int
a2 *= a2 - a4 // uint
a3 *= 0 ^ (a3 | a3) << (a2 ^ a2 - a4 ^ a2 - a2 ^ a4 << a4) // int
a3 -= (a3 ^ a3 << (0 + 3 * 2)) >> (a2 & a4) // int
v1 := a4 >> (1 ^ 2 * uint32(1) << 1) | (a2 & a4 * a2 ^ a4) >> (3 * 3 & uint32(1) << a4) ^ (a2 + a2 ^ a4) >> (a4 >> 1 ^ a4 * a2) // uint
_ = v1
return a3 ^ a3 - a3 - a3 + (int(1) << v1) << (a2 | v1) & a3 >> (3 & 2 - uint16(a5)) & a3 + a3 - a3 + a3 ^ a3
}
