package collection

import (
	"testing"
)

func TestToMap(t *testing.T) {
	students := []*Student{
		{id: 3, name: "Foo", class: "A"},
		{id: 5, name: "Beer", class: "A"},
	}
	expectedNameById := map[int]string{
		3: "Foo",
		5: "Beer",
	}

	nameById := ToMap[Student, int, string](students, func(student *Student) int {
		return student.id
	}, func(student *Student) string {
		return student.name
	})

	for id, name := range expectedNameById {
		if nameById[id] != name {
			t.Fatalf("Expceted %s - Got %s", name, nameById[id])
		}
	}
}

func TestGroupBy(t *testing.T) {
	studentFoo := &Student{id: 3, name: "Foo", class: "A"}
	studentBeer := &Student{id: 5, name: "Beer", class: "A"}
	studentLOL := &Student{id: 7, name: "LOL", class: "B"}
	students := []*Student{
		{id: 3, name: "Foo", class: "A"},
		{id: 5, name: "Beer", class: "A"},
		{id: 7, name: "LOL", class: "B"},
	}
	expectedResult := map[string][]*Student{
		"A": {studentFoo, studentBeer},
		"B": {studentLOL},
	}

	studentsByClass := GroupBy[Student](students, func(student *Student) string {
		return student.class
	})

	for class, expectedGroupedStudents := range expectedResult {
		studentsInClass := studentsByClass[class]
		for i := range studentsInClass {
			if studentsInClass[i].id != expectedGroupedStudents[i].id {
				t.Fatalf("Expceted %d - Got %d", expectedGroupedStudents[i].id, studentsInClass[i].id)
			}
		}
	}
}
