# Golang utils
## Collection utils
### Filter
#### FindOne[T any](collection []*T, condition func(element *T) bool) *T
```go
nameQuery := "Boo"
matchingStudent := FindOne[Student](students, func(student *Student) bool {
    return student.name == nameQuery
})
```

#### FindAll[T any](collection []*T, condition func(element *T) bool) []*T
```go
classQuery := "A"
matchingStudents := FindAll(students, func(student *Student) bool {
    return student.class == classQuery
})
```

### Mapping
#### Transform[T any, R any](collection []*T, transform func(*T) R) []R
```go
studentNames := Transform[Student, string](students, func(student *Student) string {
    return student.name
})
```

#### ToMap[T any, K comparable, V any](collection []*T, getKey func(*T) K, getValue func(*T) V) map[K]V
```go
nameById := ToMap[Student, int, string](students, func(student *Student) int {
    return student.id
}, func(student *Student) string {
    return student.name
})
```

#### GroupBy[T any, K comparable](collection []*T, getKey func(*T) K) map[K][]*T
```go
studentsByClass := GroupBy[Student](students, func(student *Student) string {
    return student.class
})
```

# Required
Golang 1.8 or above