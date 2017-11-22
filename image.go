package main

var a = []string{"01b.jpg", "02b.jpg", "03b.jpg", "04b.jpg", "05b.jpg", "06b.jpg", "07b.jpg", "08b.jpg", "09b.jpg", "10b.jpg", "11b.jpg", "12b.jpg", "13b.jpg", "14b.jpg", "15b.jpg", "16b.jpg", "01w.jpg", "02w.jpg", "03w.jpg", "04w.jpg", "05w.jpg", "06w.jpg", "07w.jpg", "08w.jpg", "09w.jpg", "10w.jpg", "11w.jpg", "12w.jpg", "13w.jpg", "14w.jpg", "15w.jpg", "16w.jpg"}

// removes one item from the slice
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
