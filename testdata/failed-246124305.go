// gostress -seed 246124305 -want 15
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 3, 9, ); got != 15 {
fmt.Printf("f1_ssa(2, 9, 3, ) = %v, wanted 15\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 int, a3 int8, ) int64 {
switch {} // prevent inlining
a1 += a3 * a1 // int8
a2--
v1 := (uint(0) >> ((2 << 1 - 0 * 0) >> (2 * 0 | 2))) // uint
_ = v1
v1++
a2 *= 3 * a2 | a2 // int
a3++
v1++
v2 := (2 + 2 * 0 | 3 + uint8(1) << v1 * uint8(3) << 0) << (v1 >> (2 - 1) | (v1 + v1) << (1 ^ 0)) // uint8
_ = v2
v1 = v1 * v1 - v1 << (2 | 0) * v1 + v1 * v1 | 2 | v1 // uint
a2 = a2 // int
return int64(0) >> 2 ^ 1 + 3 | int64(3) << 2 - (int64(3 ^ 0) >> (v1 - v1)) >> 2 ^ (int64(0) << v1) << v2 * (3 | 2) | (int64(2 ^ 1) >> v2)
}
