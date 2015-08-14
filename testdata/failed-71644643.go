// gostress -seed 71644643 -want 79014964751
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(3, ); got != 79014964751 {
fmt.Printf("f1_ssa(3, ) = %v, wanted 79014964751\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, ) uint64 {
switch {} // prevent inlining
v1 := ((1 ^ 3 & 0) | 0 | uint8(1)) // uint8
_ = v1
a1 -= uint64(uint(0 + 3) << (v1 << 2) - 3) - a1 * a1 ^ a1 + 2 + a1 * (a1 << a1 - 0 * a1) // uint64
a1 -= a1 >> (2 ^ 0 & 1) // uint64
a1 = a1 // uint64
v1++
a1 *= a1 * a1 - a1 ^ (a1 | a1) << 0 // uint64
v1--
v1 *= (((0 | v1) >> a1) & v1 | v1 + v1 - v1 >> a1) // uint8
return a1 * a1 >> 2 | a1 << 1 ^ a1 - a1 << (2 & 0) ^ (a1 << 3 | a1 << v1) >> ((2 | 3) << (v1 & v1))
}
