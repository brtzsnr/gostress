// gostress -seed 555113326 -want 1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 2, 9, 5, 8, ); got != 1 {
fmt.Printf("f1_ssa(6, 2, 9, 5, 8, ) = %v, wanted 1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint8, a3 int8, a4 int, a5 int8, ) int {
switch {} // prevent inlining
a5 += (a5 >> a1 ^ a3 ^ a5 | a3 << 3 * (a5 ^ a3 + a5)) << ((a2 ^ a2 | a2) - a2) // int8
a2 *= (a2 << (3 * 2)) | (a2 & a2 ^ a2) << (0 + 0 * 0) | (a2 & a2 + a2 - (a2 & a2)) >> 2 // uint8
a5 -= a3 + a5 & 2 | a3 + int8(2) >> 2 | a3 - a5 + (a5 * a3) >> 1 * a5 // int8
a1--
a1 = (a1 >> ((0 << a1 | 3 * 1) << (a1 * a1))) << (a1 & a1 & a1 | a1 * a1 + a1 | a1 * a1 << 1 & a1) // uint
a5++
a2 -= a2 - a2 * a2 ^ a2 | a2 | a2 + a2 - a2 ^ a2 * a2 * a2 * a2 & a2 ^ a2 - a2 // uint8
a2 -= ((((a2 - a2) << (a1 * a1)) >> (uint32(1) >> 1 + 2 * 2)) >> 1) // uint8
a4 -= 2 - a4 ^ a4 | 3 * (a4 + a4) ^ a4 >> 1 * (a4 >> 1 ^ a4 & a4 & a4 | a4 | a4) << (a1 & a1 << (a1 + a1)) // int
a2 *= ((a2 >> (a1 - a1 | a1)) << (uint16(3 | 2) << 2 | 0)) << (a1 & a1) // uint8
return a4 + (a4 + a4) << 0 + a4 & a4 | a4 ^ ((a4 ^ int(0) >> 1) << a1) << (1 << 2 | 2 | 1 ^ 3 >> a1 - 1 - 1)
}
