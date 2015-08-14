// gostress -seed 253927441 -want 16144481089644265473
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(8, 5, ); got != 16144481089644265473 {
fmt.Printf("f1_ssa(8, 5, ) = %v, wanted 16144481089644265473\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint64, a2 uint, ) uint {
switch {} // prevent inlining
a2 += a2 // uint
a2 = a2 << (uint32(a1) ^ (uint32(2) << 3 - 3 + 2) << (1 & 1 & uint32(2) >> a1)) // uint
a2 += a2 >> 0 ^ a2 >> 2 & a2 & a2 - a2 + a2 * a2 | a2 ^ (a2 + a2) << 3 ^ a2 // uint
a2 = a2 * a2 // uint
a1 *= (a1 << (3 * a2) - a1 - a1 - a1 >> a2 ^ a1 >> 2 + a1 ^ a1) // uint64
a2++
a1 -= (((a1 + a1) - (1 ^ a1) << (a2 ^ a2)) << (uint8(2) * uint8(1 * 2))) >> (2 ^ 0) // uint64
a2 *= a2 & a2 // uint
v1 := 0 // int
_ = v1
a1 = (a1 >> a2 | a1 ^ a1 + 0 ^ 3) >> 1 | a1 ^ a1 + a1 >> 1 - a1 & a1 + a1 << 2 & a1 // uint64
return a2
}
