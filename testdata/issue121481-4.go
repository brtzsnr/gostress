// ./a6868.go:1: doasm: notfound ft=9 tt=16 00132 (/usr/local/google/home/mosoi/source/gocode/src/github.com/brtzsnr/gostress/driver/a6868.go:9)	XORL	$648518346341351424, R14 9 16
package main
import "fmt"
func main() {
fmt.Println(f1_ssa())
}
func f1_ssa() int {
switch {} // prevent inlining
v1 := (uint32(3 + 3) << (3 + 1 - 1 + uint32(3) << 3)) >> (uint32(0) >> (3 * 0 * 3 * 2 * 3 ^ 2 & 1 ^ 1)) // uint32
v2 := (v1 ^ 2 * v1 << 2) & v1 >> 3 & v1 ^ v1 * 3 * v1 & v1 * v1 ^ v1 * v1 // uint32
v3 := v2 // uint32
v4 := v3 | v1 >> (0 + 3 + 3) // uint32
v5 := v4 + v3 - v3 * v1 ^ v1 << v1 ^ v1 << (1 >> 3) - v2 - v1 | v3 | v4 << 3 >> (v4 + v4 * v1 & v2) + v4 - (v4 * v3 >> 1) // uint32
v6 := (v5 + v5 << (2 & 2) + v4 - v4 << 3 & v4 << 0) * 3 // uint32
v7 := v6 >> 1 | v3 + v3 & v3 - v2 & v6 * v4 << v2 & v1 & v4 & v3 - v1 ^ v5 << 1 - 3 >> 3 ^ 0 + v2 // uint32
v8 := v7 ^ uint32(3 + 2 | 0 ^ 0 * 3 | 1 ^ uint64(3) >> 3 ^ 1 | 2 - 0 + uint64(2) >> 0) // uint32
v9 := v8 & v5 | v4 ^ v6 // uint32
v10 := int((0 << 2) ^ 2 - 3 << v9 & 1 << (uint8(1 ^ 2) >> (0 * 0)) | 2 & 1 ^ 2 >> v1 & 2 + 1 & 1 - 1 * 0 & 3 + 2 + 2 + 1 & 1 & 0 ^ (0 >> 0 ^ 1 | 1 >> 1 ^ 1) * (1 ^ 3 ^ 0 >> (0 - 3 | 3 + 3))) // int
_ = v10
v5--
v5--
v11 := uint(1) // uint
_ = v11
v12 := (2 + 1 - 1) & 2 - 3 * 3 | int64(0) >> 0 >> 0 << (2 | 1 + 0 >> (2 ^ 1 & 2) << (v11 >> 2 - v11 ^ v11 * v11 * v11 | 0)) + (int64(3) >> v11 >> (v11 ^ v11) | 0) ^ int64(0) << 1 // int64
_ = v12
v7--
v13 := int((1 & 3) << (uint8(3) - 1) & 0 * 1 << (0 << 1 << (uint8(2) << 0)) - 2 >> (v1 >> 3 | v3 ^ v7 | v1 + v3 & v5 - (v8 & v7) << (2 & 2 ^ 1) + v7 ^ v8 << 2 << (0 & 0) * v9 * (v3 ^ v9))) // int
_ = v13
v9--
v4++
v14 := 2 ^ (3 * 2 & 3 ^ uint64(3 + 1) << (2 | 1) - 2 - 2 & 0 - uint64(3) << 3 - 3 - 3) // uint64
_ = v14
return v10 >> (uint16(uint64(1) >> 1) >> (3 * 1) | 3 + 2 + 2 + 3 * 1 + 3 | 2 ^ 0 & 0)
}
