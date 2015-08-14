// gostress -seed 274417915 -want 1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, 5, 1, 9, 5, ); got != 1 {
fmt.Printf("f1_ssa(9, 1, 3, 6, 8, ) = %v, wanted 1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint8, a3 int8, a4 int8, a5 uint, ) int {
switch {} // prevent inlining
a5 += a5 // uint
a2 += (a2 + (a2 - a2) << (uint8(1) << 2 & 3 | a2)) // uint8
a5 -= ((a5 << (a1 & a1)) >> ((2 + 0) << 0)) >> (3 * a2 - a2 * a2 + a2) & ((a1 + a1 ^ a5 * a1) >> (uint16(1 & 0) >> (0 | 0))) << (3 ^ 3 & 0 & 2 ^ 0 + 1) // uint
a2 += ((a2 << 2 ^ a2 << a2) << (1 + 0 + 1) ^ a2 - a2 + a2 & a2) // uint8
a2++
a4 -= a4 & 2 | a4 | a4 - a3 - a3 | a4 & 3 * a4 - a3 * a3 | a4 | a4 + (a3 - a4 - a4) ^ a3 // int8
a2 = (uint8(3) << a5) >> (2 | 0 + uint16(0) << 2) + uint8(3) - a2 - 1 // uint8
a2++
a5 -= ((a1 + a5) >> (1 | 2 & 2) ^ a1 + a1) >> (uint16(0) >> (a2 >> (a5 >> 3))) // uint
v1 := a1 // uint
_ = v1
return 2 + 3 ^ 3 | 0 - int(0 * 1) << v1 - int(1) << 1 ^ int(2) >> 0 | int(1) >> v1 - 3 & 0 - 3 | 0 - 3
}
