# Практическая работа №4
## Лапин Д.С. ПИМО-01-24

Общая часть
```go
package main

import (
  "bufio"
	"errors"
	"fmt"
  "os"
	"math"
	"strings"
)
```

### Задание 1.1. Сумма цифр числа 
Напишите программу, которая принимает целое число и вычисляет сумму его цифр.
```go
func task1_1(){
  if x == 0 {
		return 0
	}
	return x%10 + task1(x/10)
}
```

### Задание 1.2
```go
func task1_2() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите температуру (например, 36.6C или 98.6F): ")
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Ошибка ввода:", err)
        return
    }

    input = strings.TrimSpace(input)

    // Парсим ввод
    var temp float64
    var scale string

    n, err := fmt.Sscanf(input, "%f%s", &temp, &scale)
    if err != nil || n != 2 {
        fmt.Println("Некорректный ввод. Пожалуйста, введите температуру в формате 36.6C или 98.6F.")
        return
    }

    scale = strings.ToUpper(scale)

    switch scale {
    case "C":
        fahrenheit := temp*9/5 + 32
        fmt.Printf("%.2f°C = %.2f°F\n", temp, fahrenheit)
    case "F":
        celsius := (temp - 32) * 5 / 9
        fmt.Printf("%.2f°F = %.2f°C\n", temp, celsius)
    default:
        fmt.Println("Неизвестная шкала. Используйте C для Цельсия или F для Фаренгейта.")
    }
}
```

### Задание 1.3. Удвоение каждого элемента массива
Напишите программу, которая принимает массив чисел и возвращает новый массив, где каждое число удвоено.

```go
func task1_3(x []int) []int {
	x[0] = x[0] * 2
	if len(x) == 1 {
		return x
	}
	x = append(x[:1], task3(x[1:])...)
	return x
}
```

### Задание 1.4. Объединение строк
Напишите программу, которая принимает несколько строк и объединяет их в одну строку через пробел.
```go
func task1_4(s []string) string {
	return strings.Join(s, " ")
}
```

### Задание 1.5. Расчет расстояния между двумя точками
Напишите программу, которая вычисляет расстояние между двумя точками в 2D пространстве.
```go
func task1_5(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}
```

## 2. Задачи с условным оператором

### Задание 2.1. Проверка на четность/нечетность
Напишите программу, которая проверяет, является ли введенное число четным или нечетным.
```go
func task2_1(x int) string {
	if x%2 == 0 {
		return "Четное"
	} else {
		return "Нечетное"
	}
}
```

### Задание 2.2. Проверка высокосного года
Напишите программу, которая проверяет, является ли введенный год високосным.
```go
func task2_2(year int) string {
	if year%400 == 0 {
		return "Високосный"
	}
	if year%100 == 0 {
		return "Невисокосный"
	}
	if year%4 == 0 {
		return "Високосный"
	}
	return "Невисокосный"
}
```

### Задание 2.3. Определение наибольшего из трех чисел
Напишите программу, которая принимает три числа и выводит наибольшее из них.
```go
func task2_3(x, y, z int) int {
	var maxNum = x
	if y > maxNum {
		maxNum = y
	}
	if z > maxNum {
		maxNum = z
	}
	return maxNum
}
```

### Задание 2.4. Категория возраста
Напишите программу, которая принимает возраст человека и выводит, к какой возрастной группе он относится (ребенок, подросток, взрослый, пожилой. В комментариях указать возрастные рамки).
```go
func task2_4(age int) {
	if age <= 12 {
		return "Ребенок"
	} else if age <= 21 {
		return "Подросток"
	} else if age <= 61 {
		return "Взрослый"
	} else {
		return "Пожилой"
	}
}
```

### Задание 2.5. Проверка делимости на 3 и 5
Напишите программу, которая проверяет, делится ли число одновременно на 3 и 5.
```go
func task2_5(inp int) string {
	if inp%3 == 0 && inp%5 == 0 {
		return "Делится"
	} else {
		return "Не делится"
	}
}
```

## 3. Задачи на циклы

### Задание 3.1 Факториал числа.
Напишите программу, которая вычисляет факториал числа.
```go
func task3_1(inp int) int {
	var acc = 1
	for inp > 1 {
		acc *= inp
		inp--
	}
	return acc
}
```

### Задание 3.2. Числа Фибоначчи
Напишите программу, которая выводит первые "n" чисел Фибоначчи.
```go
func task3_2(inp int) []int {
	if inp <= 0 {
		panic(errors.New("wrong input"))
	}
	var result = make([]int, inp)
	result[0] = 0
	if inp >= 2 {
		result[1] = 1
	}
	for i := 2; i < len(result); i++ {
		result[i] = result[i-2] + result[i-1]
	}
	return result
}
```

### Задание 3.3. Реверс массива
Напишите программу, которая переворачивает массив чисел.
```go
func task3_3(inp []int) []int {
	var result = make([]int, len(inp))
	for i := 0; i < len(inp); i++ {
		result[i] = inp[len(inp)-1-i]
	}
	return result
}
```

### Задание 3.4. Поиск простых чисел
Напишите программу, которая выводит все простые числа до заданного числа.
```go
func task3_4(inp int) []int {
	var initial_numbers = make([]int, inp-2)
	var cnt = 0
	for i := 0; i < len(initial_numbers); i++ {
		initial_numbers[i] = i + 2
	}
	for i := 0; i < len(initial_numbers); i++ {
		if initial_numbers[i] == 0 {
			continue
		}
		var divider = initial_numbers[i]
		cnt++
		for j := i + 1; j < len(initial_numbers); j++ {
			if initial_numbers[j]%divider == 0 {
				initial_numbers[j] = 0
			}
		}
	}
	var out = make([]int, cnt)
	var idx = 0
	for i := 0; i < len(initial_numbers); i++ {
		if initial_numbers[i] != 0 {
			out[idx] = initial_numbers[i]
			idx++
		}
	}
	return out
}
```

### Задание 3.5.Сумма чисел в массиве
Напишите программу, которая вычисляет сумму всех чисел в массиве.
```go
func task3_5(inp []int) int {
	var sum = 0
	for i := 0; i < len(inp); i++ {
		sum += inp[i]
	}
	return sum
}
```


### Тест
```go
func main() {
	fmt.Println(task1_1(123))
	fmt.Println(task1_2(30))
	fmt.Println(task1_3([]int{1, 2, 3, 4}))
	fmt.Println(task1_4([]string{"Hello", "go"}))
	fmt.Println(task1_5(1, 2, 3, 9))
	fmt.Println(task2_1(10))
	fmt.Println(task2_2(2024))
	fmt.Println(task2_3(2, 3, 1))
	fmt.Println(task2_4(22))
	fmt.Println(task2_5(15))
	fmt.Println(task3_1(4))
	fmt.Println(task3_2(4))
	fmt.Println(task3_3([]int{1, 2, 3}))
	fmt.Println(task3_4(30))
	fmt.Println(task3_5([]int{1, 2, 3}))
}
```
