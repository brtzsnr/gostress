// gostress -seed 3198325008 -want 3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 1, 5, 4, ); got != 3 {
fmt.Printf("f1_ssa(1, 1, 5, 4, ) = %v, wanted 3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int64, a3 uint8, a4 uint, ) int8 {
switch {} // prevent inlining
a3++
a4 = (a4 >> 0 | a4 + a4) >> (0 & 1 | 1 >> 2) * a4 & a4 ^ (a4 | a4) + a4 | a4 & a4 << 2 // uint
a4 = (a4 << 3 ^ a4 | a4 + a4 * a4 ^ a4 * a4 * a4 + a4 * a4 * a4 * a4) << ((uint32(1) >> 3) - 0 & 1 * 1 + 3) // uint
a2--
a2 = (a2 >> 3 | a2 >> 1 + a2 & a2) << ((a4 & a4) << (0 & 2) ^ a4 >> a4 - a4 >> 3) - a2 // int64
a1 = a3 | a1 - a1 - a3 << 0 ^ a3 - a1 + a1 ^ a3 & a3 + a3 // uint8
a1 *= (3 ^ a1 * a1 | a3 >> 1 ^ a3 + a1) // uint8
v1 := int8(2) >> ((0 >> a4) << (0 & 0) ^ 0 << a1 - 0) ^ 1 // int8
_ = v1
v1--
v1 = (v1 + v1 & v1 ^ v1 ^ v1) >> 2 // int8
return (v1 << a4) >> (3 * 3) | (v1 + v1 ^ v1) - v1 ^ v1 * v1 & v1 - v1 + (v1 ^ v1) << (a4 + a4) | v1 << (a4 ^ a4)
}
