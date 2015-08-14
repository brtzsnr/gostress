// gostress -seed 323812802 -want -64
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 1, 7, ); got != -64 {
fmt.Printf("f1_ssa(8, 1, 7, ) = %v, wanted -64\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 int8, a3 uint8, ) int64 {
switch {} // prevent inlining
a2++
a1 *= a1 // uint64
a3 = a3 >> (a3 - a3) * (a3 | a3) - a3 - a3 + a3 | a3 << a1 + a3 * a3 | 0 // uint8
a3++
a2 = a2 ^ a2 ^ a2 ^ (a2 ^ a2) << 3 + a2 - a2 | a2 * a2 - a2 - a2 | (a2 - a2 & a2) // int8
v1 := a1 // uint64
_ = v1
a3--
a1 *= a1 & (v1 ^ v1 << v1) >> (1 >> 0) | a1 // uint64
a3--
return (int64(3 & 0) << (a3 << 0) & 2 & 2 - int64(2 * 1) << (3 + 2)) >> ((a3 ^ a3) << (3 + 0 * 2 ^ uint16(0) << (a3 | a3)))
}
