// gostress -seed 112713334 -want 121
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 0, ); got != 121 {
fmt.Printf("f1_ssa(7, 0, ) = %v, wanted 121\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 uint8, ) int8 {
switch {} // prevent inlining
v1 := uint64(2) >> ((a2 + a2 * a2 - a2 & a2 & a2 ^ a2) << 2) // uint64
_ = v1
v2 := (int(3) << v1 + 3) >> (a2 * a2 | a2) - 1 * 2 + 0 ^ 0 + 2 * 0 | 0 + int(2) << (1 + 2 - 2 | 0 | 3 + 0) // int
_ = v2
v3 := int8(3) // int8
_ = v3
a2--
v1--
v1 *= v1 // uint64
a2 = a2 ^ (a2 >> 3) << ((a1 << 3) ^ a1 >> (3 ^ 2)) // uint8
a2 = a2 + a2 // uint8
v3 *= ((v3 * v3 ^ v3) << (0 ^ v1 * v1 ^ v1)) << (1 ^ 3) - v3 >> v1 // int8
return (v3 ^ v3) >> (1 * 1) ^ v3 ^ v3 | v3 << v1 - v3 * v3 & v3 ^ v3 ^ v3
}
