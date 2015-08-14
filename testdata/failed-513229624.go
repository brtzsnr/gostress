// gostress -seed 513229624 -want -15
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 5, 4, 0, ); got != -15 {
fmt.Printf("f1_ssa(8, 5, 4, 0, ) = %v, wanted -15\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint, a3 int, a4 int64, ) int64 {
switch {} // prevent inlining
a4--
a1 -= (a1 >> (2 >> 0) - 1 * a1 - a1 >> a1 | a1 >> 3 + a1 + a1 - a1 & a1) // uint64
a2 += a2 | (a2 + a2 * a2 ^ a2 & 3 & a2) // uint
a4 += ((a4 >> a1) << (uint16(3) << a1)) + a4 << 3 | a4 * (a4 * 1) >> (3 ^ 0) + a4 << 3 - a4 >> a2 + (a4 | a4) << (a1 << a2) | (a4 ^ a4) & 2 * a4 - a4 // int64
a3 *= (a3 >> a2 * a3 + a3 >> 1 ^ a3 & a3) >> a1 // int
a2 -= a2 // uint
a1++
a1--
return a4 << (uint8(3) << a1) * a4 ^ a4
}
