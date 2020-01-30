package main

import "fmt"

func Index(slice []int, predicate func(a int) bool) int{
	for i, value := range slice{
		if predicate(value){
			return i
		}
	}
	return -1
}

func Any(slice []int, predicate func(a int) bool) bool{
	return Index(slice, predicate) != -1
}

func Find(slice []int, predicate func(a int) bool) int{
	return slice[Index(slice, predicate)]
}

func None(slice []int, predicate func(a int) bool) bool{
	return Index(slice, predicate) == -1
}

func All(slice []int, predicate func(a int) bool) bool{
	return Index(slice, func(a int) bool {
		return !predicate(a)
	}) == -1
}

func main() {
	slice  := []int{1, 5, 10, 10, 50}
	//zet := Index(slice, func(a int) bool {
	//	return a < 0
	//})
	fine := Find(slice, func(a int) bool {
		return a < 0
	})
	fmt.Println(fine)
}
/*
Написать на базе анонимных функций и замыканий функции:
Any - true, если хотя бы один элемент из слайса удовлетворяет предикату
All - true, если все элементы из слайса удовлетворяют предикату
Index - возвращает индекс первого элемента, удовлетворяющего предикату (или -1)
Find - возвращает первый найденный элемент (подумайте, что возвращать, если не найдено - panic или error)
None - true, если ни один элемент из коллекции не удовлетворяет предикату

Any, All, None, Find должны быть построены на базе функции Index

На все функции нужны авто-тесты
 */