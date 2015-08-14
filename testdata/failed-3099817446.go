// gostress -seed 3099817446 -want 143
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 3); got != 143 {
fmt.Printf("f1_ssa(7, 3) = %v, wanted 143\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int64, a2 uint64) uint8 {
switch {} // prevent inlining
v1 := int8(0 | 2) // int8
v1--
v2 := uint8(int8(a1) - (v1 * v1) >> a2 * (v1 | v1 + v1)) // uint8
return v2 | v2 ^ v2 - v2 ^ v2 ^ v2 & v2 + v2 + v2 + v2 << a2 * v2
}
