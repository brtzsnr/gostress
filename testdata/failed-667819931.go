// gostress -seed 667819931 -want 9604
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 1, 8, 7, ); got != 9604 {
fmt.Printf("f1_ssa(1, 1, 8, 7, ) = %v, wanted 9604\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint64, a3 int64, a4 uint8, ) int {
switch {} // prevent inlining
v1 := (3 & 1 * int(2) << a2 + 3 * int(3) >> (1 & 1) ^ 0 - 3 - 1) // int
_ = v1
v1 += v1 - (v1 + v1 + v1 & v1) >> (2 ^ 1 << 1) - (v1 + v1 ^ v1) >> (a2 | a2 | a2 & a2) * v1 * (v1 * v1) >> a4 // int
v2 := (1 | a3) // int64
_ = v2
a1++
v1 *= v1 // int
v2--
v1 *= v1 // int
a2 *= a2 // uint64
a2 = (a2 ^ a2 + a2) >> 2 + a2 | (a2 << (2 & 0)) >> ((3 & 2) >> a2 + 3 << 3) // uint64
a3++
a2++
v1 -= v1 - v1 // int
v2 *= (v2 ^ v2 + v2 | a3 + a3) << (a4 << (a1 - a1) - (a4 + a4) >> (a4 - a4) ^ a4) // int64
v1 += v1 & v1 // int
a2 = a2 & a2 ^ a2 // uint64
a1 *= a1 & a1 + a1 - a1 >> (a4 & a4) & (a1 << 1) >> a4 * a1 + a1 - a1 | a1 & a1 ^ a1 - (a1 << 2) >> (3 & 1) // uint
return v1 << ((2 * 0 | 3 & 2) >> (1 - 0 + 3 + 3) ^ 1)
}
