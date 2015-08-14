// gostress -seed 155996679 -want 8861855670579224452
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(3, 3, 7, ); got != 8861855670579224452 {
fmt.Printf("f1_ssa(3, 3, 7, ) = %v, wanted 8861855670579224452\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 int64, a3 uint8, ) int64 {
switch {} // prevent inlining
a2 -= ((a2 ^ a2) >> (3 | 3)) * (a2 & a2 - a2) >> 2 & ((a2 - a2) >> (uint32(2) >> 3)) >> 3 ^ (a2 << 3) >> (1 | 0) * (a2 + a2) << a3 // int64
a1++
a2 = a2 * a2 + a2 - a2 ^ a2 + (a2 << (3 | 1)) << 2 // int64
a1 = a1 // uint
a2 -= (a2 << 1 & a2) << (2 + 1 - 0 ^ 2) - (a2 ^ a2 - a2 & a2 * a2) >> (uint64(0 ^ 1) << (1 + 3) + 3 ^ 0 * 0) // int64
a2 = a2 ^ (a2 << a3) - (a2 >> a3 ^ a2 + a2 | a2) >> (2 ^ 3 + 2 >> 0) // int64
a2 = (a2 + a2 - a2) >> ((a1 ^ a1) << (a1 & a1)) + a2 * a2 - a2 | 0 | a2 << (0 + 3) & a2 // int64
a3 -= a3 - uint8(1 ^ 0 & 2 - 1 & 1 ^ 2 & 3) // uint8
a3++
a1 -= a1 // uint
return a2
}
