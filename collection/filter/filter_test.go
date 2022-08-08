package collection

import (
	"testing"
)

type Student struct {
	id    int
	name  string
	class string
}

func TestFindOne(t *testing.T) {
	students := []*Student{
		{id: 1, name: "Foo", class: "A"},
		{id: 2, name: "Beer", class: "A"},
	}

	testCases := []struct {
		inputName         string
		expectedStudentId int
	}{
		{inputName: "Foo", expectedStudentId: 1},
		{inputName: "Beer", expectedStudentId: 2},
	}

	for i, testCase := range testCases {
		matchingStudent := FindOne[Student](students, func(student *Student) bool {
			return student.name == testCase.inputName
		})
		if matchingStudent.id != testCase.expectedStudentId {
			t.Fatalf("Case %d: Expected %d - Got %d", i, testCase.expectedStudentId, matchingStudent.id)
		}
	}
}

func TestFindAll(t *testing.T) {
	students := []*Student{
		{id: 1, name: "Foo", class: "A"},
		{id: 2, name: "Beer", class: "A"},
	}
	testCases := []struct {
		inputClass         string
		expectedStudentIds []int
	}{
		{inputClass: "A", expectedStudentIds: []int{1, 2}},
	}

	for i, testCase := range testCases {
		matchingStudents := FindAll(students, func(student *Student) bool {
			return student.class == testCase.inputClass
		})
		if len(matchingStudents) != len(testCase.expectedStudentIds) {
			t.Fatalf("Case: %d: Expected %d elements - Got %d elements", i, len(testCase.expectedStudentIds), len(matchingStudents))
		}
	}
}
