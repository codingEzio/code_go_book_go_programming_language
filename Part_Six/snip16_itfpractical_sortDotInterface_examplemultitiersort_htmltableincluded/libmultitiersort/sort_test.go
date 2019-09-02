package multitiersort

import (
	"sort"
	"testing"
)

func Cmp(a, b []Person, t *testing.T) {
	if len(a) != len(b) {
		t.Log("different lengths")
		t.Logf("%s\n%s", a, b)
		t.Fail()
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Logf("different elements, starting at %d", i)
			t.Logf("%s\n%s", a, b)
			t.Fail()
			return
		}
	}
}

func TestByColumns_Age(t *testing.T) {
	people := []Person{
		{"Alice", 12},
		{"Alice", 20},
	}
	col := &ByColumns{people, nil, 2}
	col.Select(col.LessAge)

	sort.Sort(col)
	Cmp(people, []Person{
		{"Alice", 12}, {"Alice", 20},
	}, t)
}

func TestByColumns_Name(t *testing.T) {
	people := []Person{
		{"Bob", 12},
		{"Alice", 12},
	}
	col := &ByColumns{people, nil, 2}
	col.Select(col.LessName)

	sort.Sort(col)
	Cmp(people, []Person{
		{"Alice", 12}, {"Bob", 12},
	}, t)
}

func TestByColumns_NameAge(t *testing.T) {
	people := []Person{
		{"Alice", 20},
		{"Bob", 12},
		{"Bob", 20},
		{"Alice", 12},
	}
	col := &ByColumns{people, nil, 2}
	col.Select(col.LessAge)
	col.Select(col.LessName)

	sort.Sort(col)
	Cmp(people, []Person{
		{"Alice", 12},
		{"Alice", 20},
		{"Bob", 12},
		{"Bob", 20},
	}, t)
}
func TestByColumns_AgeName(t *testing.T) {
	people := []Person{
		{"Alice", 20},
		{"Bob", 12},
		{"Bob", 20},
		{"Alice", 12},
	}
	col := &ByColumns{people, nil, 2}

	col.Select(col.LessName)
	col.Select(col.LessAge)

	sort.Sort(col)
	Cmp(people, []Person{
		{"Alice", 12},
		{"Bob", 12},
		{"Alice", 20},
		{"Bob", 20},
	}, t)
}

func TestByColumns_SumOfAgeDigitsNameAge(t *testing.T) {
	people := []Person{
		{"Aaron", 9},
		{"Aaron", 81},
		{"Alice", 20},
		{"Bob", 12},
		{"Bob", 20},
		{"Alice", 12},
	}
	maxComparisons := 3
	col := &ByColumns{people, nil, maxComparisons}

	col.Select(col.LessAge)
	col.Select(col.LessAge)
	col.Select(col.LessName)
	col.Select(col.LessSumOfAgeDigits)

	sort.Sort(col)
	Cmp(people, []Person{
		{"Alice", 20},
		{"Bob", 20},
		{"Alice", 12},
		{"Bob", 12},
		{"Aaron", 9},
		{"Aaron", 81},
	}, t)
	if len(col.columns) > maxComparisons {
		t.Errorf("Want %d comparisons, got %d", maxComparisons, len(col.columns))
	}
}
