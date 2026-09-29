package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/krystaltechz/math-calculator/internal/fractions"
	"github.com/krystaltechz/math-calculator/internal/operations"
)

var reader = bufio.NewReader(os.Stdin)

// input читает строку с консоли
func input(prompt string) string {
	fmt.Print(prompt)
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

// inputFloat читает число с консоли
func inputFloat(prompt string) float64 {
	for {
		s := strings.ReplaceAll(input(prompt), ",", ".")
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v
		}
		fmt.Println("❌ Введите корректное число!")
	}
}

// printMenu — красивое меню без уравнений
func printMenu() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════╗")
	fmt.Println("║          🧮 МАТЕМАТИЧЕСКИЙ КАЛЬКУЛЯТОР 🧮            ║")
	fmt.Println("╠══════════════════════════════════════════════════════╣")
	fmt.Println("║  АРИФМЕТИКА:                                         ║")
	fmt.Println("║    1. Сложение                                       ║")
	fmt.Println("║    2. Вычитание                                      ║")
	fmt.Println("║    3. Умножение                                      ║")
	fmt.Println("║    4. Деление                                        ║")
	fmt.Println("║    5. Возведение в степень                           ║")
	fmt.Println("║    6. Квадратный корень                              ║")
	fmt.Println("║    7. Модуль числа                                   ║")
	fmt.Println("║    8. Факториал                                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════╣")
	fmt.Println("║  ТРИГОНОМЕТРИЯ И ЛОГАРИФМЫ:                          ║")
	fmt.Println("║    9. Синус (sin)                                    ║")
	fmt.Println("║   10. Косинус (cos)                                  ║")
	fmt.Println("║   11. Тангенс (tan)                                  ║")
	fmt.Println("║   12. Натуральный логарифм (ln)                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════╣")
	fmt.Println("║  ДРОБИ И ПРОЦЕНТЫ:                                   ║")
	fmt.Println("║   13. Сложение дробей                                ║")
	fmt.Println("║   14. Вычитание дробей                               ║")
	fmt.Println("║   15. Умножение дробей                               ║")
	fmt.Println("║   16. Деление дробей                                 ║")
	fmt.Println("║   17. Процент от числа                               ║")
	fmt.Println("╠══════════════════════════════════════════════════════╣")
	fmt.Println("║    0. Выход                                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝")
}

// main — главная функция
func main() {
	for {
		printMenu()
		choice := input("👉 Ваш выбор: ")

		switch choice {
		case "0":
			fmt.Println("\n👋 До свидания!")
			return

		// ===== АРИФМЕТИКА =====
		case "1":
			fmt.Println("\n--- СЛОЖЕНИЕ ---")
			a := inputFloat("Первое число:  ")
			b := inputFloat("Второе число:  ")
			fmt.Printf("✅ %g + %g = %g\n", a, b, operations.Add(a, b))

		case "2":
			fmt.Println("\n--- ВЫЧИТАНИЕ ---")
			a := inputFloat("Уменьшаемое:   ")
			b := inputFloat("Вычитаемое:    ")
			fmt.Printf("✅ %g - %g = %g\n", a, b, operations.Sub(a, b))

		case "3":
			fmt.Println("\n--- УМНОЖЕНИЕ ---")
			a := inputFloat("Первый множитель:  ")
			b := inputFloat("Второй множитель:  ")
			fmt.Printf("✅ %g × %g = %g\n", a, b, operations.Mul(a, b))

		case "4":
			fmt.Println("\n--- ДЕЛЕНИЕ ---")
			a := inputFloat("Делимое:    ")
			b := inputFloat("Делитель:   ")
			r, err := operations.Div(a, b)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Printf("✅ %g ÷ %g = %g\n", a, b, r)
			}

		case "5":
			fmt.Println("\n--- ВОЗВЕДЕНИЕ В СТЕПЕНЬ ---")
			a := inputFloat("Основание:  ")
			b := inputFloat("Степень:    ")
			fmt.Printf("✅ %g ^ %g = %g\n", a, b, operations.Pow(a, b))

		case "6":
			fmt.Println("\n--- КВАДРАТНЫЙ КОРЕНЬ ---")
			a := inputFloat("Число:  ")
			r, err := operations.Sqrt(a)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Printf("✅ √%g = %g\n", a, r)
			}

		case "7":
			fmt.Println("\n--- МОДУЛЬ ЧИСЛА ---")
			a := inputFloat("Число:  ")
			fmt.Printf("✅ |%g| = %g\n", a, operations.Abs(a))

		case "8":
			fmt.Println("\n--- ФАКТОРИАЛ ---")
			// дебильный метод inputFloat
			// fmt.Scan()
			a := int64(inputFloat("Число:  "))
			r := operations.Factorial(a)
			fmt.Printf("✅ %d! = %d\n", a, r)

		// ===== ТРИГОНОМЕТРИЯ =====
		case "9":
			fmt.Println("\n--- СИНУС ---")
			a := inputFloat("Угол (в радианах):  ")
			fmt.Printf("✅ sin(%g) = %g\n", a, operations.Sin(a))

		case "10":
			fmt.Println("\n--- КОСИНУС ---")
			a := inputFloat("Угол (в радианах):  ")
			fmt.Printf("✅ cos(%g) = %g\n", a, operations.Cos(a))

		case "11":
			fmt.Println("\n--- ТАНГЕНС ---")
			a := inputFloat("Угол (в радианах):  ")
			fmt.Printf("✅ tan(%g) = %g\n", a, operations.Tan(a))

		case "12":
			fmt.Println("\n--- НАТУРАЛЬНЫЙ ЛОГАРИФМ ---")
			a := inputFloat("Число (> 0):  ")
			r, err := operations.Ln(a)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Printf("✅ ln(%g) = %g\n", a, r)
			}

		// ===== ДРОБИ =====
		case "13", "14", "15", "16":
			opName := map[string]string{
				"13": "СЛОЖЕНИЕ ДРОБЕЙ",
				"14": "ВЫЧИТАНИЕ ДРОБЕЙ",
				"15": "УМНОЖЕНИЕ ДРОБЕЙ",
				"16": "ДЕЛЕНИЕ ДРОБЕЙ",
			}
			fmt.Printf("\n--- %s ---\n", opName[choice])
			fmt.Println("Формат: числитель/знаменатель (например 3/4)")
			s1 := input("Первая дробь:   ")
			s2 := input("Вторая дробь:   ")

			f1, err1 := fractions.Parse(s1)
			f2, err2 := fractions.Parse(s2)
			if err1 != nil || err2 != nil {
				fmt.Println("❌ Неверный формат дроби")
				continue
			}

			var result fractions.Fraction
			var sign string
			switch choice {
			case "13":
				result, _ = fractions.Add(f1, f2)
				sign = "+"
			case "14":
				result, _ = fractions.Sub(f1, f2)
				sign = "-"
			case "15":
				result, _ = fractions.Mul(f1, f2)
				sign = "×"
			case "16":
				result, _ = fractions.Div(f1, f2)
				sign = "÷"
			}
			fmt.Printf("✅ %s %s %s = %s\n", f1, sign, f2, result)

		// ===== ПРОЦЕНТЫ =====
		case "17":
			fmt.Println("\n--- ПРОЦЕНТ ОТ ЧИСЛА ---")
			p := inputFloat("Процент (%):  ")
			n := inputFloat("Число:        ")
			fmt.Printf("✅ %g%% от %g = %g\n", p, n, operations.PercentOf(p, n))

		default:
			fmt.Println("❌ Неверный выбор. Попробуйте снова.")
		}

		// Пауза перед возвратом в меню
		// Странная пауза :|
		time.Sleep(5 * time.Millisecond)
		fmt.Print("Нажмите Enter, чтобы продолжить...")
		reader.ReadString('\n')
	}
}
