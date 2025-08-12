package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
func Keep[T any](a []T, f func(T) bool) []T {
    keep := []T{}
    for i := range a {
        if f(a[i]) {
            keep = append(keep, a[i])
        }
    }
    return keep
}

func Discard[T any](a []T, f func(T) bool) []T {
    keep := []T{}
    for i := range a {
        if !f(a[i]) {
            keep = append(keep, a[i])
        }
    }
    return keep
}