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

# Required
Golang 1.8 or above