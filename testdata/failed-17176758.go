// gostress -seed 17176758 -want 129
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, ); got != 129 {
fmt.Printf("f1_ssa(9, ) = %v, wanted 129\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, ) uint8 {
switch {} // prevent inlining
v1 := uint64(3 - 0 | 1) + 3 - 2 // uint64
_ = v1
a1 = a1 + uint(v1) // uint
v1--
v2 := uint8(1) // uint8
_ = v2
a1++
a1 = a1 + a1 | a1 + a1 - 2 ^ a1 * a1 & a1 & a1 ^ 3 & a1 * a1 | a1 * a1 ^ a1 * a1 // uint
v1++
v2 += (v2 << (uint32(0) << v1 * 1 + 3 | (1 ^ 2))) << ((v1 ^ v1) << (v2 + v2) + v1 + uint64(v2) & v1) // uint8
v3 := (v1 + v1 & v1 | (v1 + v1) >> 1) >> (uint32(3) >> 0 + uint32(0) >> v1 | 0 | 1 ^ 1 ^ 0 & 0 + 0) // uint64
_ = v3
v1--
return (v2 | v2 * v2 & 3)
}
