// gostress -seed 181289030 -want -2
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 3, ); got != -2 {
fmt.Printf("f1_ssa(8, 3, ) = %v, wanted -2\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint8, ) int8 {
switch {} // prevent inlining
a1 += (a1 | a1 << (a1 | 0 - a1)) << a2 // uint64
v1 := uint(1 * 3) >> a2 + 1 + 2 & 3 & 2 & 3 ^ ((uint(2) >> 0) >> (3 | 1)) >> (a2 & a2) // uint
_ = v1
v1--
a2 -= a2 << (v1 ^ v1 - v1 * v1) ^ a2 >> ((3 + 0) >> (a2 >> a2)) | 1 - a2 - a2 ^ 0 | a2 & a2 * a2 + (a2 - a2 ^ a2 + a2) << 1 // uint8
v1--
a2++
v2 := 0 - int64(a2 - a2) - 2 - 3 & 1 * 0 - int64(2) << (1 * 1) // int64
_ = v2
v1 = (v1 * v1 | v1 << 1) >> (v1 & 3) | v1 | ((v1 << 0 | v1 * v1) << (a1 & a1 ^ a1)) // uint
return int8(3 * (1 - 3 - 0)) >> ((uint16(v2) & 2) | ((uint16(3) >> a2) << (2 & 3)))
}
