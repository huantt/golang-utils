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

func TestUnique(t *testing.T) {
	testCases := []struct {
		inputValues    []string
		expectedValues []string
	}{
		{
			inputValues:    []string{"A", "a", "A", "B", "C"},
			expectedValues: []string{"A", "a", "B", "C"},
		},
	}

	for i, testCase := range testCases {
		uniqueValues := Unique(testCase.inputValues)
		if len(uniqueValues) != len(testCase.expectedValues) {
			t.Fatalf("Case %d: Expceted %d unqiue elements - Got %d", i, len(testCase.expectedValues), len(uniqueValues))
		}
		for k, expectedValue := range testCase.expectedValues {
			if uniqueValues[k] != expectedValue {
				t.Fatalf("Case %d: Expceted %s - Got %s", i, expectedValue, uniqueValues[k])
			}
		}
	}
}
