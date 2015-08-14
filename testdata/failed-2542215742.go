// gostress -seed 2542215742 -want 72
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 3, 6, 4, 0, ); got != 72 {
fmt.Printf("f1_ssa(2, 3, 6, 4, 0, ) = %v, wanted 72\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint, a3 int, a4 uint, a5 uint8, ) int8 {
switch {} // prevent inlining
a2 *= a2 << 2 // uint
a5 -= (a5 + a5 | a5 - a5 & a5 & a5 & a5 - a5 * a5 + a5 + (a5 * a5) << uint(0)) << ((uint16(1) >> 1 + 2 & 3) - (uint16(2 | 1) >> (3 - 0)) << (a5 << 0 - a5 + 1)) // uint8
a3 *= a1 & (a3 - a3) // int
v1 := (a1 ^ a3 | a1 * a1) // int
_ = v1
v1 = a3 // int
v1 = v1 - v1 << (3 & 1 ^ 0 - 0 & 1 | 2 + 3) // int
a3--
v1++
a1--
v1 -= 2 // int
return ((int8(2) << a5) >> a5) >> ((2 | 3) >> (1 | 2)) | (int8(2 + 2) >> (uint64(0) << a5)) << (0 >> a4 - 0) ^ int8(0 | 3 + 1) << (a4 | a4 >> a2) + 1 - 2 * 1 | 1 + 3 & 3
}
