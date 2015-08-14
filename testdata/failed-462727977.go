// gostress -seed 462727977 -want 96
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 2, 2, 0, ); got != 96 {
fmt.Printf("f1_ssa(4, 2, 2, 0, ) = %v, wanted 96\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 int64, a3 int64, a4 uint64, ) int8 {
switch {} // prevent inlining
a2--
a1++
a2--
a2++
a3 *= a2 // int64
v1 := a4 + a4 | a4 & a4 * a4 & (a4 & a4) - a4 - (a4 + a4) << (a4 + a4) // uint64
_ = v1
a2 = a3 ^ a3 | (a3 & a3 | a3) >> (uint8(0 * 0) << 2) // int64
v2 := ((a1 & a1 * a1 | a1) & a1 & a1 + a1 ^ a1 & a1 ^ a1 ^ a1 << v1) // int8
_ = v2
a4 *= uint64(0) // uint64
return (v2 - v2 - a1 >> v1 & a1 + v2 * v2 * a1)
}
