// gostress -seed 907218712 -want -1
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(3, 7, ); got != -1 {
fmt.Printf("f1_ssa(3, 7, ) = %v, wanted -1\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int8, ) int64 {
switch {} // prevent inlining
a1 -= ((a1 | a1) ^ uint8(0 & 2) & a1 & a1 * a1 * a1) >> 2 // uint8
a1 = (a1 ^ a1 & a1 ^ a1 * a1) >> (2 + a1 >> 0 | a1 & a1 << 0) // uint8
a1--
a2++
a1 -= (a1 - (a1 << 1) >> (2 ^ 2) * a1) // uint8
a1 *= (a1 - a1 - a1 << 2 | a1 >> 1) << 3 // uint8
a1--
return (int64(2 | 2 ^ 0 & 3) >> a1) & int64(2) >> (0 * uint64(3) >> a1) | 2 - int64(3) << (uint64(0) >> a1)
}
