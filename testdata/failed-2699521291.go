// gostress -seed 2699521291 -want 5
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(3, 1, 4, 0, 0, ); got != 5 {
fmt.Printf("f1_ssa(3, 1, 4, 0, 0, ) = %v, wanted 5\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int, a3 uint, a4 uint64, a5 uint8, ) int {
switch {} // prevent inlining
a4 -= a4 + a4 & 3 + a4 // uint64
a2--
a3 *= uint(3) << 0 & a3 | a3 ^ a3 + a3 * a3 + 2 * a3 & (a3 | a3 ^ a3 + a3 * a3 * a3 & a3) // uint
a4 = (a4 & a4 - a4 >> a5 ^ a4 | a4 - a4 ^ a4 & a4 + a4) // uint64
v1 := (int8(1 ^ 2 + 3 ^ 0 ^ 1 + 3) << 2) // int8
_ = v1
v1 += (v1 | v1 >> (a3 + a3 >> 0)) << ((a5 << (1 - 1) ^ (a5 & a5) << a4) << (3 >> (a3 + a3) | 0 & 0 & 2)) // int8
v1++
a3++
a2 += a1 - 1 & a2 + a2 << 0 + (a1 + a1 * a2 >> a3) ^ a1 // int
a4 -= (a4 ^ 3 & a4 - a4 ^ a4 - a4) >> (a5 + a5 | (a5 >> (0 << a5)) >> (a5 ^ a5 - a5 + a5)) // uint64
return a2 - a1 * a2 ^ a2 + (a1 >> a4 + a1 << a4) + (a2 | a1 ^ a1) << (3 - 1 & 2) + 0 + a2 - a2 * a2
}
