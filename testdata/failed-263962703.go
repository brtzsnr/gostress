// gostress -seed 263962703 -want 292874713267437568
package main
import "fmt"
import "os"
func main() {
if got := f1_ssa(9, 6, 8, 8, 0, ); got != 292874713267437568 {
fmt.Printf("f1_ssa(9, 6, 8, 8, 0, ) = %v, wanted 292874713267437568\n", got)
os.Exit(1)
}
}
func f1_ssa(a1 int8, a2 int8, a3 int64, a4 uint64, a5 int8, ) int64 {
switch {} // prevent inlining
a5 *= a5 | a5 // int8
v1 := a3 // int64
_ = v1
v1 -= ((v1 | v1) << (0 + 3 + 2 & 0) ^ v1) << ((3 + 0) << (uint32(2) << 1) ^ 2 | 0 ^ 0 + 2 & 1) // int64
v1 += (int64(3) >> 2 + v1 >> a4 + v1 & v1) >> 3 // int64
a5--
v1 -= v1 >> (1 | 1 | 1) * v1 ^ v1 & v1 ^ v1 >> 3 | (v1 + a3) << (3 | 2 + 2) ^ a3 | v1 ^ v1 | a3 ^ v1 * a3 // int64
v1 *= v1 | v1 & int64(2 + 0) + v1 | v1 // int64
a1--
a1--
return v1 & (v1 | v1 + v1 * v1) << (a4 - a4 & a4 + a4) & v1
}
