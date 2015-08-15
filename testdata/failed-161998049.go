// gostress -seed 161998049 -want -5137231320731654
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 7, 3, ); got != -5137231320731654 {
fmt.Printf("f1_ssa(9, 7, 3, ) = %v, wanted -5137231320731654\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 int, a3 uint8, ) int64 {
switch {} // prevent inlining
a2 *= (a2 * a2 ^ a2 | a2) >> 3 & (a2 * a2) >> 2 ^ a2 ^ (a2 << (uint32(2) >> 3)) << (1 | 1 | 0) // int
a2 = a2 // int
a3 -= a3 // uint8
a1++
a1 *= a1 * a1 ^ a1 - a1 * a1 * a1 - a1 // int64
a1--
a3 += (a3 >> a3) >> 1 + a3 ^ a3 * a3 | a3 // uint8
a2 *= a2 // int
a1 += a1 * 2 ^ a1 * a1 + a1 - a1 & 2 - a1 ^ a1 & a1 & a1 >> a3 * a1 * a1 * a1 // int64
a1--
return a1 | a1 | a1 & a1 << a3 * a1 * a1
}
