// gostress -seed 133666918 -want 132
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(2, 6, ); got != 132 {
fmt.Printf("f1_ssa(2, 6, ) = %v, wanted 132\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int, a2 uint8, ) int {
switch {} // prevent inlining
a1++
v1 := int64(2) // int64
_ = v1
v1 += int64(0 + 2 ^ 2 ^ uint(2) << 0) - (v1 >> 1 + v1 ^ v1 & v1 * 3 - v1) >> 2 // int64
a2--
v1--
v1--
a1 = a1 + a1 // int
v1 += ((v1 >> 1) >> 3) << 2 ^ ((v1 << 3) + v1 * v1 ^ v1) >> (uint(3) << (1 & 3) & uint(1 - 1) >> 0) // int64
a1 *= (int(2) << a2 ^ a1 - a1) >> a2 * a1 + (a1 & a1) | a1 >> 0 // int
v1 += (v1 | v1 >> 3 & v1 - v1) << (uint(1 - 3 + 2) << (1 | uint(1) >> (0 + 2))) // int64
return a1
}
