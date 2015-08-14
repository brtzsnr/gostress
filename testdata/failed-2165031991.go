// gostress -seed 2165031991 -want 99
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 3, 9, ); got != 99 {
fmt.Printf("f1_ssa(7, 3, 9, ) = %v, wanted 99\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint, a3 uint64, ) uint8 {
switch {} // prevent inlining
a3 *= a3 >> (2 + 3 - 1 + 0) & (a3 << (3 << a2) & a3 + a3 & a3) // uint64
a1 *= a1 // uint8
a2 = a2 + a2 & a2 & a2 * (a2 - a2) << a1 * a2 << 1 & a2 >> a1 ^ (a2 & a2) | a2 * a2 | a2 - a2 // uint
a1--
a2 -= (a2 - a2 & a2 - a2 >> a3 + a2) >> ((a3 | a3) >> 0) & (a2 | a2 >> 0) >> ((2 | 0 * 0 & 2) << a3) // uint
v1 := 3 & 3 * (2 ^ 2) & ((int64(0) << 3) << (a2 * a2)) << (a2 >> 1 * a2) * 0 & int64(2) >> a3 // int64
_ = v1
a3 = a3 // uint64
v1 -= (((v1 ^ v1) << a1) << a1 | v1 ^ v1 >> (a3 >> a1)) >> (((3 - 1 * 2 ^ 3) << (1 << (0 - 0))) >> ((a3 >> a1 - a3) << (1 & 2 | uint32(0) >> 1))) // int64
a2 += a2 // uint
v1 += (v1 & (v1 << 3) ^ (v1 - v1 | v1 ^ v1) << (0 + 0)) << (3 * 3 ^ 0 & 2 - 0 - 0) // int64
return ((a1 << a2 ^ a1 + a1) | uint8(1 | 1 + 0 | 2)) >> (a3 & a3 * (a3 ^ a3 & a3))
}
