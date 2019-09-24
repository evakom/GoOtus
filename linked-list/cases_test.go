package linkedlist

import "testing"

var newListTestCases = []struct {
	name      string
	in        []interface{}
	out       []interface{}
	expLength int
}{
	{
		name:      "10 items",
		in:        []interface{}{1, 2, 3, 4, 5, 6, 7, 88, 99, 1000},
		out:       []interface{}{1, 2, 3, 4, 5, 6, 7, 88, 99, 1000},
		expLength: 10,
	},
	{
		name:      "2 items",
		in:        []interface{}{1, 3},
		out:       []interface{}{1, 3},
		expLength: 2,
	},
	{
		name:      "no items",
		in:        []interface{}{},
		out:       []interface{}{},
		expLength: 0,
	},
	{
		name:      "1 item",
		in:        []interface{}{999},
		out:       []interface{}{999},
		expLength: 1,
	},
}

var RemoveTestCases = []struct {
	name        string
	in          []interface{}
	out         []interface{}
	RemoveNth   int
	expLength   int
	expectedErr error
}{
	{
		name:      "10 items, remove 5th",
		in:        []interface{}{1, 2, 3, 4, 5, 6, 7, 88, 99, 1000},
		RemoveNth: 5,
		out:       []interface{}{1, 2, 3, 4, 6, 7, 88, 99, 1000},
		expLength: 9,
	},
	{
		name:        "no items, no remove",
		in:          []interface{}{},
		RemoveNth:   1,
		out:         []interface{}{},
		expLength:   0,
		expectedErr: ErrEmptyList,
	},
	{
		name:      "1 items, remove it",
		in:        []interface{}{1000},
		RemoveNth: 1,
		out:       []interface{}{},
		expLength: 0,
	},
	{
		name:      "2 items, remove last",
		in:        []interface{}{99, 1000},
		RemoveNth: 2,
		out:       []interface{}{99},
		expLength: 1,
	},
	{
		name:      "2 items, remove first",
		in:        []interface{}{99, 1000},
		RemoveNth: 1,
		out:       []interface{}{1000},
		expLength: 1,
	},
}

var pushPopTestCases = []struct {
	name     string
	in       []interface{}
	actions  []checkedAction
	expected []interface{}
}{
	{
		name: "PushFront only",
		in:   []interface{}{},
		actions: []checkedAction{
			pushFront(4),
			pushFront(3),
			pushFront(2),
			pushFront(1),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PushBack only",
		in:   []interface{}{},
		actions: []checkedAction{
			pushBack(1),
			pushBack(2),
			pushBack(3),
			pushBack(4),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PopFront only, pop some elements",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popFront(1, nil),
			popFront(2, nil),
		},
		expected: []interface{}{3, 4},
	},
	{
		name: "PopFront only, pop till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popFront(1, nil),
			popFront(2, nil),
			popFront(3, nil),
			popFront(4, nil),
			popFront(nil, ErrEmptyList),
		},
		expected: []interface{}{},
	},
	{
		name: "PopBack only, pop some elements",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popBack(4, nil),
			popBack(3, nil),
		},
		expected: []interface{}{1, 2},
	},
	{
		name: "PopBack only, pop till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			popBack(4, nil),
			popBack(3, nil),
			popBack(2, nil),
			popBack(1, nil),
			popBack(nil, ErrEmptyList),
		},
		expected: []interface{}{},
	},
	{
		name: "mixed actions",
		in:   []interface{}{2, 3},
		actions: []checkedAction{
			pushFront(1),
			pushBack(4),
			popFront(1, nil),
			popFront(2, nil),
			popBack(4, nil),
			popBack(3, nil),
			popBack(nil, ErrEmptyList),
			popFront(nil, ErrEmptyList),
			pushFront(8),
			pushBack(7),
			pushFront(9),
			pushBack(6),
		},
		expected: []interface{}{9, 8, 7, 6},
	},
}

// checkedAction calls a function of the linked list and (possibly) checks the result
type checkedAction func(*testing.T, *List)

func pushFront(arg interface{}) checkedAction {
	return func(t *testing.T, l *List) {
		l.PushFront(arg)
	}
}

func pushBack(arg interface{}) checkedAction {
	return func(t *testing.T, l *List) {
		l.PushBack(arg)
	}
}

func popFront(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, l *List) {
		v, err := l.PopFront()
		if err != expectedErr {
			t.Errorf("PopFront() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopFront() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func popBack(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *List) {
		v, err := ll.PopBack()
		if err != expectedErr {
			t.Errorf("PopBack() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopBack() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}
