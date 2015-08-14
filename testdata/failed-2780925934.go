// gostress -seed 2780925934 -want 0
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 9, 1, 6, 3, ); got != 0 {
fmt.Printf("f1_ssa(2, 9, 1, 6, 3, ) = %v, wanted 0\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 int64, a3 uint8, a4 int8, a5 int, ) int {
switch {} // prevent inlining
v1 := (a5 - a5) << (1 | 2) & a5 | a5 ^ a5 - a5 ^ a5 << 2 ^ a5 - a5 & a5 * a5 & a5 & a5 | a5 | a5 & a5 | a5 | (1 + a5) // int
_ = v1
v1 = ((a5 * v1) << (a3 >> 2) & v1 - v1 ^ a5 + a5 | (v1 * v1 << a1) << (1 ^ 0 ^ 0)) >> a1 // int
a4--
v1 *= v1 - v1 + (v1 & v1 & v1 + v1 ^ v1 * v1) >> ((1 + 2) * 2) // int
v1--
a5++
v1--
a4--
return v1
}
