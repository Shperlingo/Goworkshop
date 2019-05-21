package main

import "testing"

func TestT05(t *testing.T) {
	tables := []struct {
		a   int
		b   int
		res int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{1, 1, 4},
	}

	for _, table := range tables {
		res := DoSum(table.a, table.b)
		if res != table.res {
			t.Errorf("%d+%d expect %d got: %d", table.a, table.b, res, table.res)
		}
	}
}
