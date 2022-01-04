package main

import "sort"

// UniqueIntsWithMap returns the unique int slice.
func UniqueIntsWithMap(slice []int) []int {
	var result []int
	uniqueMap := make(map[int]bool)
	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = true
			result = append(result, value)
		}
	}
	return result
}

// UniqueIntsWithSort2 returns the unique int slice.
func UniqueIntsWithSort2(slice []int) []int {
	if len(slice) == 0 {
		return nil
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	var result []int
	for i := 0; i < len(slice); i++ {
		if i+1 < len(slice) && slice[i] == slice[i+1] {
			continue
		}
		result = append(result, slice[i])
	}
	return result
}

// UniqueIntsWithSort returns the unique int slice.
func UniqueIntsWithSort(slice []int) []int {
	if len(slice) == 0 {
		return nil
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	var result []int
	result = append(result, slice[0])
	for i := 1; i < len(slice); i++ {
		if slice[i-1] == slice[i] {
			continue
		}
		result = append(result, slice[i])
	}
	return result
}
