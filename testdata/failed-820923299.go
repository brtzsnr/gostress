// gostress -seed 820923299 -want -84
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 5, 8, 4, 5, ); got != -84 {
fmt.Printf("f1_ssa(9, 5, 8, 4, 5, ) = %v, wanted -84\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint8, a3 int8, a4 int64, a5 uint64, ) int8 {
switch {} // prevent inlining
a3++
a5 *= a5 ^ a5 // uint64
a3--
a1++
a4 = (a4 + a4 >> 2) << ((uint32(1 | 1) >> (a2 | a2)) << ((a2 << 1) << (3 + 2)) ^ 0 + 2 & 3 - 2 - 1 | 0) // int64
a3 -= a3 // int8
a3 += a3 | a3 + a3 & a3 + a3 | 3 // int8
a4++
a3 += a3 >> a2 // int8
return ((int8(3) >> a5 - a3 << a2 * (a3 * a3) >> (3 ^ 1)) << ((a1 * a1 | a1) >> a1)) << uint(2)
}
