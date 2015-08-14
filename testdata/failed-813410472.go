// gostress -seed 813410472 -want 20
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 1, ); got != 20 {
fmt.Printf("f1_ssa(4, 1, ) = %v, wanted 20\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint8, ) uint64 {
switch {} // prevent inlining
v1 := uint(2 + 0) >> (0 | 0) & 2 ^ 1 & 0 & (0 * 2 - 3) - (2 + uint(1) << 1 | 1 ^ 0 | uint(1) << a2) << (3 & 2 ^ 3 * 2 + 2) // uint
_ = v1
v1--
v2 := 2 & 0 * 2 & 3 - 0 * 3 + (int(2 & 0) << v1 - int(3) >> (a1 & a1)) << ((2 & 3 * 1 - 2) << 1) // int
_ = v2
v1++
v2 = v2 * v2 << 0 // int
a2 *= a2 | a2 + a2 * a2 - (a2 | a2) >> 2 & (a2 - a2) | a2 * (a2 | a2 * a2) >> (1 >> 0 * 2 << a2) // uint8
v1--
v1 = v1 - (v1 * v1 | v1 | v1 + v1) << (2 | 1 ^ 2) // uint
v1 = v1 + v1 - v1 - v1 & (v1 | v1) ^ uint(a1) ^ v1 - v1 & v1 // uint
return a1 | ((a1 << v1 + a1 * a1) >> ((v1 << 0) << v1)) >> 0
}
