// gostress -seed 2345929641 -want 659712
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 9, 1, 2, ); got != 659712 {
fmt.Printf("f1_ssa(7, 9, 1, 2, ) = %v, wanted 659712\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint64, a3 int8, a4 int, ) int {
switch {} // prevent inlining
a1 *= ((a1 & a1) >> (a2 - a2)) ^ a1 | (a1 * a1) >> 0 ^ a1 // uint8
v1 := a2 // uint64
_ = v1
v1 -= v1 - v1 // uint64
a3 += a3 >> (a1 | a1 & a1 + a1 - 2 | a1 | a1 + a1 - a1 ^ 2 ^ a1 | a1 ^ a1) // int8
a1 = a1 - a1 | a1 ^ a1 * a1 + a1 - a1 | a1 * a1 // uint8
v1 = v1 // uint64
a3++
a4 += (a4 - a4 * a4 - 1 - a4) >> v1 & (a4 * 2 - a4 | a4 ^ a4 * a4 * a4) << v1 // int
v2 := v1 // uint64
_ = v2
v2 = ((v1 >> 2) - v1 >> a2 + v2 | 3 | v2 * v1 + v2 * v1 >> 2 + v1) << ((v2 << 3 - v2 << v2) >> 2 | v1 ^ v2) // uint64
return (a4 << 1 - a4) << (a1 >> (a1 * a1)) - (a4 << a1) >> (3 & 1) + a4 ^ a4 + a4 & ((a4 + a4 | a4 | a4) << (2 | 3 ^ 1 * 1)) >> (3 * 0)
}
