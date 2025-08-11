package lesson5

import "fmt"

func main() {
	var (
		a int
		b int
		c string
	)

	for {

		fmt.Printf("\n Добро пожаловать! Нажмите Enter, чтобы продолжить, или введите что угодно для выхода.")
		_, err := fmt.Scanln(&c)
		if c != "" {
			break
		}
		for {
			fmt.Println("Введите первое целое число :")
			_, err = fmt.Scan(&a)
			if err != nil {
				fmt.Println("Ошибка. Пожалуйста, вводите целое число.")
				continue
			}
			fmt.Println("Введите второе целое число: ")
			_, err = fmt.Scan(&b)
			if err != nil {
				fmt.Println("Ошибка. Пожалуйста, вводите целое число.")
				continue
			}
			result, err := divide(a, b)
			if err != nil {
				fmt.Println(err)
				break
			}

			switch {
			case result >= 10:
				fmt.Println("Результат большой")
			case result > 1 && result < 10:
				fmt.Println("Результат средний")
			default:

				fmt.Println("Результат маленький или ноль")

			}
			break
		}
	}
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Не выдумывай — делить на ноль нельзя.")
	}
	return a / b, nil

}
