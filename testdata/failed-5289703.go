// gostress -seed 5289703 -want 18446744073709550431
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 3, 0, 3, ); got != 18446744073709550431 {
fmt.Printf("f1_ssa(9, 3, 0, 3, ) = %v, wanted 18446744073709550431\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 int64, a3 uint64, a4 uint8, ) uint {
switch {} // prevent inlining
a3--
a4 += a4 // uint8
a1 += a1 - a1 + a1 * (a1 ^ a1 * a1 * a1) >> (a3 >> (a4 - a4)) // int
a1++
v1 := 3 * a2 & a2 * a2 & a2 ^ a2 & a2 * a2 & a2 & a2 ^ a2 | a2 + a2 - a2 + (a2 >> a4) << (3 >> 2) // int64
_ = v1
v1 += v1 << 1 + v1 & a2 | v1 | v1 - v1 + ((v1 & v1) << (2 - 1)) << a4 + (int64(0) >> (1 >> a4) & 2) >> (uint16(v1) >> 3 & 3 * 1 ^ 3 * 2) // int64
a4 = a4 // uint8
v2 := a2 >> (3 * 0) * v1 >> (a3 << 0) | v1 & 0 - v1 - a2 >> a4 + v1 >> (0 & 3 & 3 | 0 * uint(1 + 0) << 3) // int64
_ = v2
return uint(2) >> a3 * 2 & 2 - (3 * uint(v1))
}
