package binary

func Total(depth int) int {
	if depth == 0 {
		return 1
	}

	return (2 << depth) - 1
}

func ParentIdx(idx int) int {
	return (idx - 1) / 2
}

func LeftChildIdx(idx int) int {
	return 2*idx + 1
}

func RightChildIdx(idx int) int {
	return 2*idx + 2
}
