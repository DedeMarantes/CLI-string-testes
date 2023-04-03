package functions

import "sort"

func ReverseString(in string) string {
	var result string = ""
	for _, e := range in {
		result = string(e) + result
	}
	return result
}

func SortStringList(lista []string, inverse bool) []string {
	sort.Strings(lista)
	if inverse {
		for i, j := 0, len(lista)-1; i < j; i, j = i+1, j-1 {
			//inverte ordem da lista
			lista[i], lista[j] = lista[j], lista[i]
		}
	}
	return lista
}
