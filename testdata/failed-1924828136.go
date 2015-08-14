// gostress -seed 1924828136 -want 3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 5, ); got != 3 {
fmt.Printf("f1_ssa(2, 5, ) = %v, wanted 3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint, ) int8 {
switch {} // prevent inlining
a2--
a1++
a1++
a1++
a2 = a2 & a2 | a2 | a2 & a2 - a2 | a2 | a2 | a2 // uint
a1++
a1 -= a1 | a1 // uint8
a2 -= a2 + a2 * a2 & a2 + a2 * a2 - a2 - a2 & a2 - a2 // uint
a2 = ((a2 + a2 << a1 + a2 >> a2) << (a2 >> 2 - a2 & a2 * a2 - a2 & a2)) << ((a1 & a1) >> (a2 << a2) + a1 - a1 ^ a1 * (a1 >> 2) << 1 - a1 - a1) // uint
a2++
return (int8(0) << 2 * 3 | 2 * 1 + 1 & 0 - (int8(3) >> a2) | int8(3) << a2 ^ 0 * 1) >> (1 - uint32(0) << a2 * uint32(3 & 0) << 0 * uint32(1 + 2 + 2) >> (a2 << (a1 - a1)))
}
