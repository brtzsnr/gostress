// gostress -seed 1028317872 -want 336064488
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 0, 3, 7, 4, ); got != 336064488 {
fmt.Printf("f1_ssa(2, 0, 1, 5, 6, ) = %v, wanted 336064488\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint64, a3 uint8, a4 int64, a5 int, ) uint64 {
switch {} // prevent inlining
v1 := (((a2 | a1) >> (0 | 2) & 0 - a1 ^ a2) >> 3 & a2) // uint64
_ = v1
a1--
v1++
v1--
v1 = (a2 + a1 * v1 * v1) >> (v1 << 2 | a2 << 3) | v1 * v1 + a1 + v1 | v1 | a1 * v1 | v1 << (2 + 1 * 0) ^ a1 | 0 ^ a2 | 0 | a2 + v1 * v1 + v1 | v1 & v1 + a1 + v1 + ((v1 - v1) >> (v1 >> 2)) << 0 // uint64
v1 -= ((v1 ^ v1) ^ v1) & a1 >> 2 * a1 * v1 & v1 << (3 & 3) * v1 ^ v1 | v1 // uint64
a2 *= v1 // uint64
a3++
a5 = a5 << (2 | 1 | 0) - a5 ^ a5 - a5 + a5 ^ a5 - a5 << 2 | int(v1 | v1 - 2 * v1) ^ a5 // int
a5--
a1 = (v1 << 1 & v1 - v1) << 3 ^ (v1 & a1) | (v1 | v1 ^ v1 - v1 - v1 & a2) >> (3 ^ 3 << 0 | 1 << 0) ^ v1 // uint64
a5 *= a5 // int
v1 += (uint64((1 | 3)) - v1 & v1 | v1) << uint(a4) // uint64
a2 -= uint64((a5 >> (2 - 2)) >> 2 * a5 ^ (a5 ^ a5 ^ a5 * a5 & a5) >> (3 - 1 + 2 + 3 ^ 1 << (uint(0) >> 3))) // uint64
v1 -= a2 // uint64
v2 := a4 // int64
_ = v2
v1 -= v1 & v1 + v1 & v1 * a2 * a2 - a1 + a2 - a2 ^ v1 & v1 & ((v1 & a2) ^ v1 ^ v1 | a1) << a3 & v1 * a1 & a2 * v1 >> (1 ^ 3) - v1 & v1 >> a3 ^ a2 | (a2 | a1 - a2 << 2 * 1 | v1 + a1 | v1) >> (1 * 1 + 3 << a3) // uint64
v2++
return v1
}
