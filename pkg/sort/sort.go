package sort

type Sorter interface {
	Sort(src string) string
}

type SorterFunc func(string) string

func (m SorterFunc) Sort(src string) string {
	return m(src)
}

func New() SorterFunc {
	return Merge
}

func Merge(src string) string {
	return string(mergeSort([]byte(src)))
}

func mergeSort(slice []byte) []byte {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return merge(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
}

func merge(left, right []byte) []byte {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]byte, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
