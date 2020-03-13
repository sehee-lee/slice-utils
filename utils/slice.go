package utils

import "reflect"

func IsSlice(slice interface{}) (isSlice bool) {
	refSlice := reflect.ValueOf(slice)
	if refSlice.Kind() != reflect.Slice {
		return false
	}
	return true
}

func IsNotSlice(slice interface{}) (isNotSlice bool) {
	return !IsSlice(slice)
}

func EmptySlice(slice interface{}) (isEmpty bool) {
	refSlice := reflect.ValueOf(slice)
	if IsNotSlice(slice) {
		return true
	} else if refSlice.Len() == 0 {
		return true
	}
	return false
}

func NotEmptySlice(slice interface{}) (isNotEmpty bool) {
	return !EmptySlice(slice)
}

func ContainElement(slice interface{}, element interface{}) (exist bool, idx int) {
	if IsNotSlice(slice) {
		return false, -1
	}
	refSlice := reflect.ValueOf(slice)
	refElement := reflect.ValueOf(element)

	if refElement.Kind() == reflect.Invalid {
		return false, -1
	}

	if refSlice.Len() == 0 || refSlice.Elem().Kind() != refElement.Kind() {
		return false, -1
	}

	for i := 0; i < refSlice.Len(); i++ {
		if refSlice.Index(i).Interface() == refElement.Interface() {
			return true, i
		}
	}

	return false, -1
}

func RemoveItem(slice interface{}, idx int) (removed bool) {
	if EmptySlice(slice) {
		return false
	}
	refSlice := reflect.ValueOf(slice).Elem()

	result := reflect.AppendSlice(refSlice.Slice(0, idx), refSlice.Slice(idx+1, refSlice.Len()))
	refSlice.Set(result)
	return true
}