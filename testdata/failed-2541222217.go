// gostress -seed 2541222217 -want 3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 3, 4, ); got != 3 {
fmt.Printf("f1_ssa(8, 3, 4, ) = %v, wanted 3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 int, a3 int, ) uint {
switch {} // prevent inlining
v1 := uint(1 + 2) // uint
_ = v1
a1 = a1 - (a1 + a1) - (a1 + a1) >> 3 - a1 << 3 * 3 | a1 | (a1 - a1 ^ a1 & (a1 >> v1)) << (uint(2) << 2 | v1 | v1) // int8
v1 = v1 << 0 // uint
v2 := v1 // uint
_ = v2
a1++
v3 := 0 & 1 * 3 * 2 | uint64(2 ^ 3) << (v1 * v1) + uint64(1) >> 1 ^ 0 & 0 ^ 3 * 2 * 0 + 3 * 1 * 2 + uint64(2) >> (3 ^ 3) ^ (uint64(2) << v1) // uint64
_ = v3
v4 := uint8((((v3 + v3) << 0) << ((v2 << 3) >> uint64(v1))) >> (v3 >> 0 & v3 >> v3 ^ v3 & v3 - 0 ^ v3)) // uint8
_ = v4
a3 -= (a3 ^ a2 >> (2 * 3)) // int
v3--
a2 -= a3 // int
return v2 << v4 * (v1 >> (0 ^ 1)) ^ v2
}
