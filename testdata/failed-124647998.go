// gostress -seed 124647998 -want -703
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 9, 1, 6, 2, ); got != -703 {
fmt.Printf("f1_ssa(1, 9, 1, 6, 2, ) = %v, wanted -703\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 int8, a3 int, a4 uint, a5 uint64, ) int64 {
switch {} // prevent inlining
a1 += a1 + a2 - a1 & a1 ^ a1 ^ (a2 * a1) + a1 + a2 - a2 + a2 & a1 + a2 >> (uint16(0) >> 0) // int8
v1 := a5 // uint64
_ = v1
v1 += v1 & v1 & v1 & v1 + v1 & v1 >> 3 - a5 | v1 | v1 ^ v1 + a5 + v1 // uint64
v1 *= (0 * v1) // uint64
a5 = (v1 << (3 & 2 & 2 - 2) - 2 ^ (v1 >> 2) >> 3) // uint64
v2 := a2 * a1 << (1 & 0 + 1) ^ a1 + a1 << a4 - a1 ^ a2 // int8
_ = v2
v1 -= a5 >> ((a4 * a4) << (0 & 2 | 0)) // uint64
v1 = a5 // uint64
a5++
v3 := a4 + a4 + a4 << 1 ^ a4 & a4 // uint
_ = v3
return 1 ^ ((int64(1) >> a4) << 1) - (int64(3) << (3 + 1)) << (2 - 0 | 2) * 3 | 3 ^ 2 ^ 2 & int64(3) << a4
}
