
// 1.Функции
//Напишите функцию, которая принимает два числа и возвращает их сумму.
func sum(num1 int, num2 int) int {
return num1 + num2
}
// Создайте функцию с несколькими возвращаемыми значениями: результат деления и остаток.
func divide(num1 int, num2 int) (int, int) {
div := num1 / num2
mod := num1 % num2
return div, mod
}
// Реализуйте функцию, которая принимает функцию как аргумент и вызывает её.
func doublingDivide(divide func() int) int {
return divide() * 2
}
//2. Массивы и срезы
// Создайте массив из 5 целых чисел и выведите их на экран
arr := [...]int{1, 2, 3, 4, 5}
fmt.Println(arr)


// Напишите функцию, которая принимает срез и возвращает его копию с удвоенными элементами.
func doubleSlice(arr []int) []int {
arr2 := make([]int, len(arr))
for i, v := range arr {
arr2[i] = v * 2
return arr2
}
}
// Реализуйте объединение двух срезов с помощью append
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{1, 2, 3, 4, 5}
fmt.Println(append(slice2, slice1...))

//3. Карты (map)
//Создайте map[string]int для хранения возраста людей по имени и выведите данные
m := make(map[string]int)
m["Alex"] = 29
m["Svetlana"] = 25
m["Nikita"] = 28
fmt.Println(m)

// Реализуйте функцию, которая подсчитывает количество вхождений каждого слова в срезе строк.
func countWords(slice []string, name string) int {
         var count int
         for _, v := range slice3 {
             if v == name {
             count++
			 }
         }
     return count
}
//Удалите ключ из карты и проверьте его наличие
delete(m, "Alex")
v, e := m["Alex"]
if e {
fmt.Println("Нашли:", v)
} else {
fmt.Println("Нет такого ключа")
}
// 4.Структуры и методы
// Определите структуру Person с полями Name и Salary
type Person struct {
Name   string
Salary int
}
}

