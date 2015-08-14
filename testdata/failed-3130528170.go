// gostress -seed 3130528170 -want 125
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 1, 7, 7, 9, ); got != 125 {
fmt.Printf("f1_ssa(2, 1, 7, 7, 9, ) = %v, wanted 125\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint8, a3 int8, a4 int8, a5 int8, ) int8 {
switch {} // prevent inlining
a3++
a5 = a4 * a5 - a4 * a5 * a5 | a4 + a3 * a3 | a4 - (a5 ^ a5) >> (0 & 0) * a4 & a3 | a3 | a5 ^ a3 // int8
a4 = ((a3 + a4 * a3) << ((0 << a2) << (3 >> 3 + 0))) << (1 >> (uint(0 | 3 - 0) << (1 + 3 - uint(0)))) // int8
a5 *= (a5 << 1) & (a5 >> 3) & a3 << (2 | 1 - 0) & (a3 & a3 + a3 << 0) * a3 // int8
a4--
a2--
a3 -= ((a3 + a3) | (a4 + a4) << 0) << ((uint32(2 + 1) << (1 ^ 0)) >> (2 | 3 - 3 | 0)) ^ (a5 ^ a4) << (uint16(3 | 0) >> (0 & 0)) * (a3 - a3) >> 3 + a4 & a5 - a3 // int8
a5 -= a5 // int8
return a3 & ((a3 ^ a3) >> (3 * 2)) - (a4 >> a2 * a4 | 2) << a2
}
