// gostress -seed 1718211778 -want -28
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(7, 7, 5, 3, 2, ); got != -28 {
fmt.Printf("f1_ssa(7, 7, 5, 3, 2, ) = %v, wanted -28\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint, a2 int8, a3 int8, a4 uint8, a5 uint64, ) int8 {
switch {} // prevent inlining
a2++
a1 = a1 + a1 - ((a1 << a5 - a1 ^ a1) >> ((3 * 2) << (0 >> a5))) << (3 & 2) // uint
a5 -= a5 + a5 // uint64
a4--
a5 = ((a5 * 3) ^ a5 | a5 - a5 + (a5 - a5) >> (0 + 1 ^ 2)) << (a1 | a1 << 3 ^ a1 ^ a1 >> a1 - (a1 | 2) >> (3 | 0 & 1)) // uint64
v1 := a3 << a5 | a2 - (a3 ^ a3) | (a3 & a2) ^ a2 >> a5 ^ a3 ^ a3 * (a2 >> a5) >> (1 ^ 2 ^ 0 - 2) // int8
_ = v1
v1++
v1 *= v1 // int8
a1 -= a1 // uint
return v1 << (1 + 3 - 2 & 2)
}
