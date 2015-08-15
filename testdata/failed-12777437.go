// gostress -seed 12777437 -want -5299
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 5, 0, ); got != -5299 {
fmt.Printf("f1_ssa(9, 5, 0, ) = %v, wanted -5299\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 uint8, a2 uint, a3 uint8, ) int64 {
switch {} // prevent inlining
v1 := ((int64(3 & 2) << (a2 ^ a2) * 2 & 3 ^ 3) << ((0 << (1 * 2)) << ((a1 ^ a3) << (3 & 3)))) * 1 // int64
_ = v1
a3 *= a3 ^ (a3 & a1 + a1 ^ 1) << (1 >> a1 & 0) & a3 ^ a1 - 1 ^ a3 - a1 * a1 ^ a1 - a1 // uint8
v1 += v1 | (v1 ^ v1) << (0 << 1) * v1 | v1 + v1 + v1 - v1 * v1 << 1 // int64
a2--
a3 *= a1 // uint8
v1 += (int64(3) + (v1 | v1 + v1)) << a2 // int64
a1++
v1 = v1 // int64
a2 += (a2 << 0 ^ a2 & a2 << a2) >> (1 | 3 * 3 & (1 - 0) >> (2 & 1)) - a2 ^ a2 ^ a2 + a2 & a2 << a2 ^ a2 | (a2 * a2 >> 0) << a2 + (0 + a2) >> (a2 * a2) | a2 - a2 ^ a2 | a2 & a2 - a2 >> a2 + a2 << 0 * a2 & a2 * (a2 + a2 + a2 ^ a2) * a2 * a2 & a2 & a2 | a2 // uint
return (v1 + v1 * (v1 & v1 | v1 & v1) >> (0 ^ 2)) ^ int64((a2 | a2 >> a1)) | (v1 ^ v1) << a2 - v1 >> (3 - 1) ^ v1
}
