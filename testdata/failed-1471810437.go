// gostress -seed 1471810437 -want 16769
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(1, 1, 7, 7, ); got != 16769 {
fmt.Printf("f1_ssa(1, 1, 7, 7, ) = %v, wanted 16769\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint8, a3 uint, a4 uint, ) int64 {
switch {} // prevent inlining
v1 := (int64(3) << 1 - 0 | 2 ^ 2 | 0 * int64(1) << 2 | int64(3) << a2) >> a2 // int64
_ = v1
a1--
v1--
v1 += (((v1 * v1) << (2 | 3 | 3 & 2)) << (a2 - a2 | a2 + a2)) // int64
v2 := int8(1) // int8
_ = v2
v2++
v1--
v3 := a2 << 2 // uint8
_ = v3
v2--
a4 *= a3 << (2 - 3 ^ 1 + 3 | 0 * 1 * 0) & a3 // uint
return (v1 * v1) - (v1 >> 2) << uint32(2) & v1 * v1 >> (v3 + v3) * v1 & v1 | v1
}
