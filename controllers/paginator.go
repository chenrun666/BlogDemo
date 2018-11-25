package controllers

func HandlePrevious(data int) int{
	pageIndex := data -1
	return pageIndex
}

func HandleNext(data int) int {
	pageIndex := data + 1
	return pageIndex
}