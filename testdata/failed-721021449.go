// gostress -seed 721021449 -want 11
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 9, ); got != 11 {
fmt.Printf("f1_ssa(2, 9, ) = %v, wanted 11\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint, ) int {
switch {} // prevent inlining
a2 -= (a2 + a2 & a2 | a2 - a2 + a2 - a2) << 0 - a2 * a2 - a2 - a2 & ((a2 & 2) << (0 & 3)) // uint
a1--
a2 = (uint(3 * 1 + 2) - a2) >> (uint8(3) & 1) // uint
a2 = (a2 ^ (a2 ^ a2 | a2 | a2) << 2) << (a2 + a2 ^ a2 + (a2 | a2 - a2) >> (a2 + a2 * a2 * a2)) // uint
a2++
a1 = (0 | a1) >> (2 * 3 * 1 + 3 * 2 & 2) + a1 & a1 | a1 ^ a1 - (a1 - a1) & a1 << (0 & 3 ^ 3) // int8
a2++
a1 -= ((a1 & a1 - a1) >> (1 + uint64(1) >> 2) * a1) >> (2 * 0 | 1 | 3 ^ 3 >> a2 - (2 ^ 0) << 2) // int8
return 3 + int(1 + 0 + 1) >> (a2 | a2) + int(2) << (a2 & a2) - 2 + 1 | 2 & 3 * 0 | 1 & 1
}
