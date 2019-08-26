package slave

/* Example One */

type ProfileExportable struct {
	Name string
	Age  int
}
type profileNotExportable struct {
	Name string
	Age  int
}

func ExportLowercaseProfileNotExportable(name string, age int) profileNotExportable {
	return profileNotExportable{
		Name: name,
		Age:  age,
	}
}

/* Example Two */

type Dog struct {
	Name string
	age  int
}

/* Example Three */

// If it's `Animal`, everything will be fine and without any fuss
type animal struct {
	Name string
	Age  int
}

type Cat struct {
	animal
	Cuteness int
}

/* Example Three+ */

type Animal struct {
	Name string
	Age  int
}

type Cat2 struct {
	Animal
	Cuteness int
}
