// Code generated by "stringer -type=TryAllowedResult"; DO NOT EDIT.

package sac

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Deny-1]
	_ = x[Allow-2]
}

const _TryAllowedResult_name = "UnknownDenyAllow"

var _TryAllowedResult_index = [...]uint8{0, 7, 11, 16}

func (i TryAllowedResult) String() string {
	if i < 0 || i >= TryAllowedResult(len(_TryAllowedResult_index)-1) {
		return "TryAllowedResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TryAllowedResult_name[_TryAllowedResult_index[i]:_TryAllowedResult_index[i+1]]
}
