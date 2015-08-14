// gostress -seed 1130331263 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 9, 2, 6, 6, ); got != 0 {
fmt.Printf("f1_ssa(2, 9, 2, 6, 6, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 int8, a3 int64, a4 uint8, a5 uint64, ) int8 {
switch {} // prevent inlining
a5 = a1 * a1 | a1 - a1 // uint64
a5--
a1--
a1 = a1 // uint64
a2 -= 0 | a2 | a2 | a2 * a2 * (a2 + a2) ^ a2 - a2 | a2 * 3 & a2 << 3 - a2 - a2 + a2 - (a2 >> 1) << a4 // int8
a4 += a4 // uint8
a4 = a4 >> (uint16(3) << 2 | 2 | 1) ^ a4 ^ a4 - a4 // uint8
a3--
a4 *= a4 // uint8
return a2 & (a2 >> a5) << ((a5 & a5) << 3 | a1 * a5 - a5 ^ a5)
}
