// gostress -seed 255956923 -want -3904
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 0, 1, 8, ); got != -3904 {
fmt.Printf("f1_ssa(6, 0, 1, 8, ) = %v, wanted -3904\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint8, a3 int8, a4 int64, ) int64 {
switch {} // prevent inlining
a4 *= a4 // int64
a1++
a1 -= a1 + int8((a4 * a4) << (2 - 2)) - a3 >> 2 + a3 & a3 // int8
a4--
a4++
a1--
a4 -= (((a4 & a4) << (0 << 0)) << (1 ^ 2 ^ 3) * a4 + (a4 + a4)) // int64
a4 -= ((a4 | a4) - a4 << 0 * a4 & (a4 | a4) << a2 + a4 | a4 + a4) // int64
a3 *= a1 - a3 // int8
return a4 << (1 + 0) + a4 >> a2 - a4 & a4 & a4 >> a2 - a4 & a4 & a4 >> a2 * a4 | 1 & a4 | a4 * a4 ^ (a4 & a4) >> (a2 + a2)
}
