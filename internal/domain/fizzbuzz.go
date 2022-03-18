package domain

import (
	"fmt"
	"strconv"
)

type Fizzbuzz struct {
	// TODO: add upper bound limit to avoid crashing service

	Int1  int    `form:"int1" validate:"gt=0"`
	Int2  int    `form:"int2" validate:"gt=0"`
	Limit int    `form:"limit" validate:"gt=0"`
	Str1  string `form:"str1" validate:"required"`
	Str2  string `form:"str2" validate:"required"`
}

func (f Fizzbuzz) ToString() string {
	return fmt.Sprintf("%d,%d,%d,%s,%s", f.Int1, f.Int2, f.Limit, f.Str1, f.Str2)
}

func (f Fizzbuzz) ComputeSequence() []string {
	result := make([]string, 0, f.Limit)
	// TODO: opitmize code, add memoization.

	for i := 1; i <= f.Limit; i++ {
		switch {
		case i%f.Int1 == 0 && i%f.Int2 == 0:
			result = append(result, f.Str1+f.Str2)
		case i%f.Int1 == 0:
			result = append(result, f.Str1)
		case i%f.Int2 == 0:
			result = append(result, f.Str2)
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
