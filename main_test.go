package main

import "testing"

func TestContainerFull(t *testing.T) {
	type testDataFull struct {
		con  Container
		full bool
	}

	dataItems := []testDataFull{
		testDataFull{
			Container{Capacity: 10, Quantity: 10},
			true,
		},
		testDataFull{
			Container{Capacity: 10, Quantity: 5},
			false,
		},
		testDataFull{
			Container{Capacity: 10, Quantity: 0},
			false,
		},
	}

	for _, item := range dataItems {
		result := item.con.isFull()
		if result != item.full {
			t.Errorf("FAILED, expected %v but got value '%v'", item.full, result)
		} else {
			t.Logf("PASSED, expected %v and got value '%v'", item.full, result)
		}
	}
}

func TestContainerEmpty(t *testing.T) {
	type testDataEmpty struct {
		con   Container
		empty bool
	}

	dataItems := []testDataEmpty{
		testDataEmpty{
			Container{Capacity: 10, Quantity: 0},
			true,
		},
		testDataEmpty{
			Container{Capacity: 10, Quantity: 10},
			false,
		},
		testDataEmpty{
			Container{Capacity: 10, Quantity: 8},
			false,
		},
	}

	for _, item := range dataItems {
		result := item.con.isEmpty()
		if result != item.empty {
			t.Errorf("FAILED, expected %v but got value '%v'", item.empty, result)
		} else {
			t.Logf("PASSED, expected %v and got value '%v'", item.empty, result)
		}
	}
}

func TestContainerFill(t *testing.T) {
	type testDataFill struct {
		con    Container
		expQty uint64
	}

	dataItems := []testDataFill{
		testDataFill{
			Container{Capacity: 10, Quantity: 0},
			1,
		},
		testDataFill{
			Container{Capacity: 10, Quantity: 10},
			10,
		},
		testDataFill{
			Container{Capacity: 10, Quantity: 8},
			9,
		},
	}

	for _, item := range dataItems {
		item.con.fill()
		if item.con.Quantity != item.expQty {
			t.Errorf("FAILED, expected %v but got value '%v'", item.expQty, item.con.Quantity)
		} else {
			t.Logf("PASSED, expected %v and got value '%v'", item.expQty, item.con.Quantity)
		}
	}
}
