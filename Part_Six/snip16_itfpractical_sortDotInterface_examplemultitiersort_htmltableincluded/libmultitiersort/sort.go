package multitiersort

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type ByColumns struct {
	person     []Person
	columns    []columnCmp
	maxColumns int
}

type comparison int
type columnCmp func(a, b *Person) comparison

const (
	lt comparison = iota
	eq
	gt
)

// -------------------------------------------------------

func (person Person) String() string {
	return fmt.Sprintf("%s: %d", person.Name, person.Age)
}

// -------------------------------------------------------

func (col *ByColumns) Len() int      { return len(col.person) }
func (col *ByColumns) Swap(i, j int) { col.person[i], col.person[j] = col.person[j], col.person[i] }
func (col *ByColumns) Less(i, j int) bool {
	for _, f := range col.columns {
		cmp := f(&col.person[i], &col.person[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

func (col *ByColumns) LessName(a, b *Person) comparison {
	switch {
	case a.Name == b.Name:
		return eq
	case a.Name < b.Name:
		return lt
	default:
		return gt
	}
}
func (col *ByColumns) LessAge(a, b *Person) comparison {
	switch {
	case a.Age == b.Age:
		return eq
	case a.Age < b.Age:
		return lt
	default:
		return gt
	}
}

// -------------------------------------------------------

func NewByColumns(p []Person, maxColumns int) *ByColumns {
	return &ByColumns{p, nil, maxColumns}
}

func (col *ByColumns) Select(cmp columnCmp) {
	// prepend the new comparison
	col.columns = append([]columnCmp{cmp}, col.columns...)

	// not letting the slice of comparisons grow without bound
	if len(col.columns) > col.maxColumns {
		col.columns = col.columns[:col.maxColumns]
	}
}

// -------------------------------------------------------

func sumOfDigits(n int) int {
	sum := 0
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	return sum
}

func (col *ByColumns) LessSumOfAgeDigits(a, b *Person) comparison {
	aSum := sumOfDigits(a.Age)
	bSum := sumOfDigits(b.Age)

	switch {
	case aSum == bSum:
		return eq
	case aSum < bSum:
		return lt
	default:
		return gt
	}
}
