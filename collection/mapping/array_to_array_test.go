package collection

import "testing"

type Student struct {
	id    int
	name  string
	class string
}

func TestTransform(t *testing.T) {
	students := []*Student{
		{id: 1, name: "Foo", class: "A"},
		{id: 2, name: "Beer", class: "A"},
	}
	expectedResult := []string{
		"Foo", "Beer",
	}
	studentNames := Transform[Student, string](students, func(student *Student) string {
		return student.name
	})

	for i, name := range studentNames {
		if expectedResult[i] != name {
			t.Fatalf("Expceted %s - Got %s", expectedResult[i], name)
		}
	}

}
