// Code generated by "stringer -type Fruit ./enum/fruit.go"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Apple-1]
	_ = x[Banana-2]
	_ = x[Orange-3]
}

const _Fruit_name = "AppleBananaOrange"

var _Fruit_index = [...]uint8{0, 5, 11, 17}

func (i Fruit) String() string {
	i -= 1
	if i < 0 || i >= Fruit(len(_Fruit_index)-1) {
		return "Fruit(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Fruit_name[_Fruit_index[i]:_Fruit_index[i+1]]
}
