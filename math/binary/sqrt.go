// sqrt.go
// description: Square root calculation
// details:
// Calculating the square root using binary operations and a magic number 0x5f3759df [See more](https://en.wikipedia.org/wiki/Fast_inverse_square_root)
// author(s) [red_byte](https://github.com/i-redbyte)
// time complexity: O(1)
// space complexity: O(1)
// see sqrt_test.go

package binary

func Sqrt(n float32) float32 { return 1 / FastInverseSqrt(n) }
