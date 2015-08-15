// gostress -seed 1783714943 -want 191
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(0, 4, 9, ); got != 191 {
fmt.Printf("f1_ssa(0, 4, 9, ) = %v, wanted 191\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 int, a3 uint, ) int64 {
switch {} // prevent inlining
a1 += a1 // int64
a2 += (a2 ^ a2 | a2 + a2 | a2 + a2) | (a2 >> 0) >> (2 - 0 + 2) | a2 * a2 << 3 * a2 ^ a2 ^ (a2 ^ a2 ^ a2 >> 3 + a2 ^ a2 ^ a2 + a2) // int
v1 := uint64(2) // uint64
_ = v1
a3--
v1 = v1 ^ v1 ^ v1 >> (3 - 0 - 2) ^ v1 & v1 ^ (v1 ^ v1 - v1 << 3) >> ((0 >> 0) >> (1 + 0)) - v1 + v1 - v1 // uint64
v1++
v2 := int8(0) // int8
_ = v2
v3 := 0 * a2 - ((a2 ^ a2) << v1) >> ((v1 + v1) << (0 + 1 ^ 0)) ^ a2 | a2 ^ a2 - a2 | a2 ^ a2 | a2 | a2 // int
_ = v3
a3++
return (a1 - a1 + a1 ^ a1 >> a3 - a1 & a1) | a1 + (int64(3) << a3 - 1 + a1) >> (0 ^ 3) ^ a1 - (a1 << v1 * a1 * a1) << (0 ^ 3 & 3) * a1 & a1 ^ a1 & a1 & a1
}
