// gostress -seed 2387713120 -want 18874441
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 6, ); got != 18874441 {
fmt.Printf("f1_ssa(6, 6, ) = %v, wanted 18874441\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 uint, ) int {
switch {} // prevent inlining
a2 = a2 // uint
a1--
a1 += (a1 * a1 | a1) << (2 | uint64((3 * 0)) >> ((uint16(3) >> 3) >> a2)) // int8
v1 := int64(2) << 0 // int64
_ = v1
v1--
a2 += (a2 >> (1 + 3 | 1 - 1)) << (uint8(3 ^ 3 | 3) << a2) ^ ((a2 >> 2) << 2) << (0 & 3 * 3 ^ uint16(1) << 1) // uint
v1 = (v1 | v1 & v1) << ((a2 * 3 & a2) << 0 | a2 * a2 << 3 | a2 & a2 >> a2) // int64
v1++
v1 += v1 // int64
return (int(3) >> a2 | 0 + 3 * int(3 | 1) << a2 + int(3) << 3 * int(3) << 0 * 2 | 0 ^ 1 + 2) >> (1 + 0 * uint16(0) >> 1 | uint16(0 | 2) << 2 & 1 ^ 2 & 0 & 3)
}
