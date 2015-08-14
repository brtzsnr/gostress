// gostress -seed 896431477 -want 8
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(4, 0, ); got != 8 {
fmt.Printf("f1_ssa(4, 0, ) = %v, wanted 8\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint64, ) uint64 {
switch {} // prevent inlining
a2++
a2 *= a2 << (2 | 2 * 0 & 0 - 0 | 2 - uint8(0) << a2 - 0 * 0 | 2 * uint8(2) >> a2) // uint64
a1 += a1 >> 2 & (a1 << 3 - int(2) | (a1 + a1)) << (a2 << 0 & a2 - a2 | a2 & a2 + a2 | a2) // int
a1 *= (a1 ^ a1 | a1) | a1 + a1 * a1 >> 2 ^ a1 >> 2 // int
a2 = a2 // uint64
a1++
v1 := int8(a1) & 1 - 2 | 3 & int8(3) << 3 // int8
_ = v1
v1++
return a2 + a2 | a2 * a2 | a2 ^ a2 | a2 ^ a2 * a2 ^ a2
}
