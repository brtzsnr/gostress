// gostress -seed 87036070 -want -964062
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 0, ); got != -964062 {
fmt.Printf("f1_ssa(1, 0, ) = %v, wanted -964062\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint8, ) int {
switch {} // prevent inlining
v1 := 1 & int64(3 & 3 & 2) >> (1 + 1 | 2 | 2) | (int64(2) >> a2 ^ int64(2) << 0) << (2 - uint32(2) >> 3 | 2 + 3) // int64
_ = v1
a1++
a1 += (a1 - 2 + a1 & a1 * a1) >> 2 + a1 - a1 | a1 & a1 + a1 * a1 + int(3) // int
a2--
a1--
a1 *= a1 // int
a2--
a2 -= a2 + a2 // uint8
v2 := (v1 | v1 ^ v1 * v1 * v1) & v1 | v1 + v1 & v1 & v1 << 2 // int64
_ = v2
a1 = a1 >> (uint32(2 - 2) << 0) ^ a1 + a1 ^ a1 * a1 - a1 ^ 1 - a1 * a1 // int
return a1 >> a2 + (int(3) << a2) << (3 * 2) + a1 | a1 * a1 & a1 * a1 + a1
}
