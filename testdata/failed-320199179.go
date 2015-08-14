// gostress -seed 320199179 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, ); got != 0 {
fmt.Printf("f1_ssa(6, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, ) uint {
switch {} // prevent inlining
a1 = a1 ^ ((a1 >> 3) << (uint64(0) >> 1)) << (3 - 3 ^ uint8(3) << 1) ^ int8(0) ^ a1 << 2 ^ a1 + a1 ^ a1 | a1 * a1 // int8
v1 := a1 >> ((3 - 2 + 2 - uint64(1 | 3) << (3 ^ 3)) << 2) // int8
_ = v1
a1++
v2 := 1 & 0 + 3 * 2 & uint8(3) << 2 * 3 - 2 * 2 | 3 | 0 | 2 * 2 * 3 // uint8
_ = v2
v3 := uint64(2) // uint64
_ = v3
a1 -= (v1 & 3 - a1 | v1 + a1 - a1) * a1 & a1 * a1 ^ a1 // int8
a1 -= v1 + v1 // int8
v4 := v2 >> ((3 >> v3 - 0 >> 1) << 3 * 3 | 0 * 1 ^ 2 | 1 + 2 | 1) // uint8
_ = v4
v5 := v3 // uint64
_ = v5
return (1 | uint(2 | 1) << (v3 >> v3) + 1 + 2 & 2 - 1) >> (v4 << (v5 | v5 ^ v5) * v4 | v2 ^ v2 - v2 + v2 * v2)
}
