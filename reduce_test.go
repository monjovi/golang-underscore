package underscore

import (
	"testing"
)

func TestReduce(t *testing.T) {
	v, err := Reduce([]int{ 1, 2 }, func (memo []int, n, _ int) ([]int, error) {
		memo = append(memo, n)
		memo = append(memo, n + 10)
		return memo, nil
	}, make([]int, 0))
	if err != nil {
		t.Error(err)
	}

	res, ok := v.([]int)
	if !(ok && len(res) == 4) {
		t.Error("wrong length")
	}

	if !(res[0] == 1 && res[1] == 11 && res[2] == 2 && res[3] == 12) {
		t.Error("wrong result")
	}
}

func TestChain_Reduce(t *testing.T) {
	v, err := Chain([]int{ 1, 2 }).Reduce(func (memo []int, n, _ int) ([]int, error) {
		memo = append(memo, n)
		memo = append(memo, n + 10)
		return memo, nil
	}, make([]int, 0)).Value()	
	if err != nil {
		t.Error(err)
	}

	res, ok := v.([]int)
	if !(ok && len(res) == 4) {
		t.Error("wrong length")
	}

	if !(res[0] == 1 && res[1] == 11 && res[2] == 2 && res[3] == 12) {
		t.Error("wrong result")
	}
}