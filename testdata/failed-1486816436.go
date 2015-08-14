// gostress -seed 1486816436 -want -506
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 1, ); got != -506 {
fmt.Printf("f1_ssa(8, 1, ) = %v, wanted -506\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint, ) int64 {
switch {} // prevent inlining
v1 := a1 // uint8
_ = v1
a1 -= v1 - v1 & (a1 * a1) >> 3 + 1 + a1 + v1 & (a1 * v1) << (a1 >> 0) | a1 // uint8
v2 := (int64(a2 >> v1 - a2 | a2) ^ int64(0) >> (a2 ^ a2)) << (a1 >> 2 | a1 | a1 >> (a1 + 0) ^ v1 | a1 ^ v1 << a1 - a1 - a1 * a1 << 3) // int64
_ = v2
a2++
v3 := 1 // int
_ = v3
v2++
a1 *= a1 // uint8
a2++
a2 += (a2 + a2 * a2 & a2 + (a2 * a2) << (3 ^ 1)) >> (2 | 1) + (a2 & a2) ^ a2 * a2 // uint
return v2 * v2 + v2 ^ v2 + 1 | v2 + v2
}
