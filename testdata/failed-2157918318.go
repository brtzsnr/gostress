// gostress -seed 2157918318 -want 152474225066
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 3, 2, 6, ); got != 152474225066 {
fmt.Printf("f1_ssa(9, 3, 2, 6, ) = %v, wanted 152474225066\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int, a3 uint, a4 int8, ) int64 {
switch {} // prevent inlining
a2--
a3 = (a3 & a3 | a3) << (0 - 0) | a3 - a3 + a3 ^ a3 + a3 ^ a3 * a3 >> (2 - 0) // uint
v1 := a4 // int8
_ = v1
a4 += (((v1 << 2) * v1) >> (1 & 1 * 2 & uint16(2 * 2) << (2 - 1))) >> (2 | 1) // int8
v1--
a3 += (a3 | (a3 << a3 + a3 - a3) << (1 * uint64(3) >> 3)) // uint
v1 *= a4 * v1 | v1 ^ v1 | v1 * (a4 << 2) << (2 + 0) - v1 * v1 * v1 | v1 & v1 // int8
v1 *= v1 - v1 | v1 * v1 | v1 | v1 >> 0 ^ a4 // int8
a2 -= a1 // int
v2 := (a1 << (1 + 1 ^ 1 << 2)) >> 0 - (a1 << 0 | a2 << 2) >> (2 ^ uint8(3) >> 1) & a2 << 0 & a2 & a1 - (a2 | a1) >> (0 ^ 3) // int
_ = v2
a2 *= a2 & v2 - a1 + a1 >> (3 + 1) - v2 << 1 & a1 | a2 * a2 << 0 ^ 3 & v2 - (1 - a1 * a1) // int
a1++
a2 = (a1 << 2) | a2 | a1 | a1 + (v2 - v2) >> 0 ^ ((a2 >> a3 - a2) >> (0 << (0 * 3))) // int
a2 -= ((v2 >> 0 * v2 >> 1 + (v2 & a1) << (a3 | a3)) >> (2 + 3)) // int
a1--
v3 := (v1 >> 2 - v1) >> uint64(2) // int8
_ = v3
a1 *= (a1 >> (2 ^ 3) * v2 | 3 + v2 * a2 & (a1 + v2) << (uint64(1) >> 0)) >> (1 * 3 ^ 0 & 3 ^ 0 - 2 | 2 - 3 + 2 ^ 1) // int
v2 -= a2 >> 1 ^ a2 + v2 - (a2 + a1) ^ a2 - v2 - a1 * a2 | (v2 >> 0) * a2 & a2 | a2 ^ a1 * a2 * a1 | a2 // int
v2--
v3++
return 3 & 2 & 0 ^ 1 + 1 & 1 - int64(a1 & v2 | a2 * v2)
}
