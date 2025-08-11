//1. Объявить переменную с вашим именем и вывести её.
var Alexandr string
x := Alexandr
// 2. Создать константу с числом Пи и вывести.
 const pi = 3.1415926
 fmt.Println(x)
// 3. Написать программу, которая сравнивает два числа и выводит большее.
package main
import "fmt"
func main() {
	var (
		num1 int = 15
		num2 int = 17
	)
	fmt.Println(Comparison(num1, num2))
}
	func Comparison (num1 int, num2 int) int {
		if num1 > num2 {
			return num1

		} else if num1 < num2 {}
return num2
}
}
//4.  Сделать функцию, которая проверяет, чётное ли число.
func  evenNumber (a int)string{
	if a % 2 == 0{
		return  "Четное"
	}else{
		return "Нечетное"
	}
}
//5. Использовать if-else для определения взрослый человек или нет по возрасту.
var age  int = 12

if age <=18 {
	fmt.Println( "Юный")
}else if age >18&&40 {
    fmt.Println("Взрослый")
}else{
    fmt.Println("Пожилой")
}
//6. Написать цикл for, который выводит числа от 1 до 10.
for i := 1 ; i<=10; i++ {
 fmt.Println(i)
}
//7. Использовать for с условием (while-подобный) для подсчёта суммы чисел от 1 до n.
n := 10
i := 1
s := 0
for  i<=n {
	s += i
	i++
	fmt.Println("Cумма от 1 до",n,"=",s)
}
//8. Использовать switch для вывода названия дня недели по номеру.
week := 1
switch week {
case 1:
	fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
case 4:
    fmt.Println("Thursday")
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
	fmt.Println("Sunday")
}
//9. Сделать функцию с двумя возвращаемыми значениями — результат и ошибка (например, деление).
func sum(a int ,b int ) (int,error){
	var s int = a+b
	if s >= 100{
		return s,errors.New("too")
}
	return s,nil
}
//10. Использовать именованные возвращаемые параметры в простой функции.
func sum(a int ,b int ) (d int,err error){
d = a+b
err = errors.New("Много")
if d >= 100{
return d,err
}
return d,nil
}
//11. Написать программу, которая пропускает вывод чётных чисел с помощью continue.

for i := 0 ; i <= 100;i++ {
	if i % 2 == 0{
	continue
}else{
		fmt.Println("Нечетное",i)
}
}

//12. Использовать break для выхода из цикла при достижении числа 5.
var m int = 0
for  i:=0; i<100; i++{
	m += i

	if m == 5 {
		break
}
}
//13. Написать программу, которая считывает число с консоли и выводит его.
package main
import "fmt"
func main() {
var a int
fmt.Scan(&a)
fmt.Println(a)
}
//14. Использовать fmt.Printf для форматированного вывода строки и числа.
var age int = 15
var name string ="Petr"
fmt.Printf("Имя: %q,Возраст: %d\n",name,age)
//15. Объявить несколько переменных в одной строке и вывести их сумму.
var a,b,c int = 1,2,3
sum :=a+b+c
fmt.Println(sum)
//16. Написать функцию, которая принимает два числа и возвращает их сумму.
func sum(a int, b int) ( int ){
	return a+b
}
//17. Создать переменную bool и вывести её значение в зависимости от условия.
var s bool
var num int = 15
if num <=15 {
	s = true
}else {
	s = false
}
fmt.Println(s)

//18. Сделать программу, которая выводит сообщение, если число отрицательное, ноль или положительное (if-else).
package main

import "fmt"

func main() {

fmt.Println(num(0))
}
func num(i int) string {
if i < 0 {
return "Отрицательное"
} else if i == 0 {
return "ноль"
} else {
return "Положительное"
}
}

//19. Использовать if с коротким объявлением переменной.
if x := 4; x <= 4 {
fmt.Println(x + 1)
} else {
fmt.Println(x)
}
}

//20. Написать простую обработку ошибки через if err != nil

result, err := divide(10, 0)
if err != nil {
fmt.Println(err)
}
fmt.Println(result)
func divide(a, b int) (int, error) {
if b == 0 {
return 0, fmt.Errorf("Не выдумывай — делить на ноль нельзя.")
}
return a / b, nil

}