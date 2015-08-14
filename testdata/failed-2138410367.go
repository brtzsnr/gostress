// gostress -seed 2138410367 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 2, 9, 0, 3, ); got != 0 {
fmt.Printf("f1_ssa(6, 2, 9, 0, 3, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 int, a3 int8, a4 uint64, a5 uint8, ) int8 {
switch {} // prevent inlining
a4++
a5--
a1 *= (a1 | 2 & a3 & a1) * int8(3 + a2) + (a3 ^ a3) << (a5 & a5) + (int8(1) >> 2) >> (0 ^ 1) ^ a1 // int8
a1++
a3 *= (a1 & a1) << (2 | 3) // int8
a3 -= a1 // int8
a1 = a3 // int8
a4 += a4 // uint64
a1 += ((a1 - a3) >> (2 ^ 3 ^ 2 & 0) - a3 - a1 >> 0 - a1) >> ((3 ^ 0) << (0 ^ 2) | (1 >> a4) >> uint32(a2) * 0) // int8
return a3 & a1 - a3 | a3 ^ a3 | a3 - a3 + a1 & a1 * (a1 - a1 | a1 - (a3 >> a5) << (uint16(2) >> 3)) >> (3 * 1 + 2 - 0)
}
