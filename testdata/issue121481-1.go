// ./a31018.go:1: doasm: notfound ft=9 tt=13 00124 (/usr/local/google/home/mosoi/source/gocode/src/github.com/brtzsnr/gostress/driver/a31018.go:11)	ORW	$22898193250, AX 9 13
package main
import "fmt"
func main() {
fmt.Println(f1_ssa())
}
func f1_ssa() uint8 {
switch {} // prevent inlining
v1 := uint16(389)
v2 := v1 + v1 | v1 * v1 * v1 | v1 & 3 | v1 * (v1 ^ 2) & v1 - v1 // uint16
v3 := ((int8(3) >> v2) & int8(0) >> v2 >> 1 & 2 ^ 1) << 1 // int8
v4 := v3 + 2 ^ v3 << 3 * 3 & v3 * 0 - v3 << v2 << v2 & v3 - v3 | int8(2) >> (v2 - v2 | v1 * v2) * 1 >> (v1 & v2 + v2 + v1) & (int8(1) >> 3 & 2 * v3 + v3) >> v2 // int8
v5 := v4 | v3 * v3 // int8
v6 := v5 * v3 // int8
v7 := ((v6 ^ v6 * v3) ^ int8(2) * (v4 >> 0)) >> 2 // int8
v9 := uint(v7 >> (v2 * v1)) ^ uint(0 + 0) >> (3 & 2) + 2 ^ 0 & 1 - 0 * uint(1) >> 1 - 3 * 0 // uint
v11 := (uint64(1) >> v9)-1 // uint64
v12 := uint64(0) >> 1 | v11 - v11 & 3 | v11 - v11 - v11 - 1 | 2 - 1 // uint64
return 0 ^ 1 + 0 - (uint8(0) >> 3 ^ uint8(1 & 3) >> v12 ^ 1 ^ 0 * 2 - 0 ^ uint8(0 * 2) >> (uint8(3) >> 0))
}
