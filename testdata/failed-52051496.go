// gostress -seed 52051496 -want -1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 0, 7, ); got != -1 {
fmt.Printf("f1_ssa(1, 0, 7, ) = %v, wanted -1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint, a3 int, ) int8 {
switch {} // prevent inlining
a3 -= a3 // int
a1++
a1 *= (a1 | a1) << a2 - (a1 + 3) << 1 // int64
a1++
a3++
a2 *= a2 + a2 | a2 + a2 | a2 * a2 - a2 & a2 ^ a2 - a2 ^ a2 ^ a2 ^ a2 + 0 - (a2 ^ a2) ^ a2 ^ a2 | a2 & a2 | a2 ^ a2 + a2 // uint
a1 -= a1 * a1 // int64
a1--
a2++
return int8(0) >> (1 << 0) - 3 ^ 1 * 3 - (int8(2) >> a2) ^ 0 ^ 1 * 1 - 0 * (int8(2 | 1) >> 1 | 2) + 2 ^ int8(1) << a2 ^ 3 | 0
}
