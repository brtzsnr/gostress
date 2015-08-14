// gostress -seed 192507284 -want 3
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(3, 2, 2, ); got != 3 {
fmt.Printf("f1_ssa(3, 2, 2, ) = %v, wanted 3\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 int8, a3 uint64, ) int8 {
switch {} // prevent inlining
a3 -= uint64(int64(3) << (0 ^ 3 + 1) - 2 + 2 & 2 - int64(0 - 2) >> (1 + 2)) // uint64
a2 *= a2 >> a1 ^ ((a2 - a2) >> (a3 - a3)) | a2 >> (2 & 2) + a2 - a2 + a2 ^ (a2 >> a3 | a2 >> 3) >> (a3 - a3 - a3) // int8
a2++
a1 -= (a1 + a1 ^ a1) // uint8
a1--
a2--
a2++
v1 := 1 & uint(0) >> (1 & 1) + 1 & 0 * 2 * 2 // uint
_ = v1
v1++
a2 -= a2 << ((a3 >> a1) - a3 >> (a3 - a3 * a3 - a3)) // int8
return a2 * a2 & a2 ^ (a2 - a2) << a3 ^ int8(1) | a2 - a2 | a2 ^ a2 | 2 ^ a2 & a2 | a2 ^ a2 - a2
}
