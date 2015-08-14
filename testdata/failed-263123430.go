// gostress -seed 263123430 -want 60
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(6, 6, ); got != 60 {
fmt.Printf("f1_ssa(6, 6, ) = %v, wanted 60\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint, ) uint64 {
switch {} // prevent inlining
a1 = (a1 - int(3) - a1 - a1 + a1 * a1 + a1 ^ a1 - a1) >> (a2 << (1 + 2 ^ 1 & (3 | 3))) // int
a2--
v1 := ((int64(1) >> a2) >> a2 ^ 2 * 1 - 3 & 0) + 0 // int64
_ = v1
v1++
a1--
v1 += v1 // int64
a2 *= a2 // uint
v1--
v1 += (v1 >> 3) // int64
return uint64(a2 << a2) >> a2 | uint64(3 | 3) << (uint64(2) << 1) + 3 | 3 | 1 ^ 0 | uint64(1) << 3 & 1 | 3 ^ 3
}
