package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
type intSlice []int
func (a intSlice) Less(i, j int) bool{
	return a[i] < a[j]
}
func (a intSlice) Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}
func (a intSlice) Len() int{
	return len(a)
}
func Solution(A []int) int {
	// write your code in Go 1.4
	res:=1
	sort.Sort(intSlice(A))
	for jj := range A{
		if A[jj] == res{
			res++
		}
	}
	return res
}

