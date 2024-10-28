## Практическая работа №5
### Лапин Д.С. ПИМО-01-24

```go
package main

import (
	"errors"
  "fmt"
  "time"
  "os"
	"math"
  "math/rand"
	"strings"
)
```

Задание 1. Проверка на простоту
   Напишите функцию, которая проверяет, является ли переданное число простым. Ваша программа должна использовать циклы для проверки делителей, и если число не является простым, выводить первый найденный делитель.
```go
func IsPrime(n int) bool {
    if n <= 1 {
        fmt.Printf("%d не является простым числом.\n", n)
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            fmt.Printf("%d не является простым числом. Первый делитель: %d\n", n, i)
            return false
        }
    }
    fmt.Printf("%d является простым числом.\n", n)
    return true
}

func main() {
    var number int
    fmt.Print("Введите число: ")
    fmt.Scanln(&number)
    IsPrime(number)
}
```

Задание 2. Наибольший общий делитель (НОД)
   Напишите программу для нахождения наибольшего общего делителя (НОД) двух чисел с использованием алгоритма Евклида. Используйте цикл `for` для вычислений.
```go
func gcd(a int, b int) int {
 for b != 0 {
  a, b = b, a%b
 }
 return a
}

func main() {
 var num1, num2 int

 fmt.Print("Введите первое число: ")
 fmt.Scan(&num1)
 fmt.Print("Введите второе число: ")
 fmt.Scan(&num2)

 result := gcd(num1, num2)
 fmt.Printf("Наибольший общий делитель чисел %d и %d равен %d", num1, num2, result)
}
```

Задание 3. Сортировка пузырьком 
   Реализуйте сортировку пузырьком для списка целых чисел. Программа должна выполнять сортировку на месте и выводить каждый шаг изменения массива.
```go
func bubbleSort(arr []int) {
 n := len(arr)
 for i := 0; i < n; i++ {
  for j := 0; j < n-i-1; j++ {
   if arr[j] > arr[j+1] {
    // Меняем местами если элемент больше следующего
    arr[j], arr[j+1] = arr[j+1], arr[j]
    // Выводим состояние массива на каждом шаге
    fmt.Println(arr)
   }
  }
 }
}

func main() {
 arr := []int{64, 34, 25, 12, 22, 11, 90}
 fmt.Println("Исходный массив:", arr)
 bubbleSort(arr)
 fmt.Println("Отсортированный массив:", arr)
}

```

Задание 4. Таблица умножения в формате матрицы
   Напишите программу, которая выводит таблицу умножения в формате матрицы 10x10. Используйте циклы для генерации строк и столбцов.
```go
func main() {
 const size = 10
 for i := 1; i <= size; i++ {
  for j := 1; j <= size; j++ {
   fmt.Printf("%d*%d=%d | ", j, i, j*i)
  }
  fmt.Println()
 }
}
```

Задание 5. Фибоначчи с мемоизацией
   Напишите функцию для вычисления числа Фибоначчи с использованием мемоизации (сохранение ранее вычисленных результатов). Программа должна использовать рекурсию и условные операторы.
```go
func fibonacci(n int, memo map[int]int) int {
 if val, exists := memo[n]; exists {
  return val
 }

 if n <= 0 {
  return 0
 } else if n == 1 {
  return 1
 }

 memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
 return memo[n]
}

func main() {
 memo := make(map[int]int)
 var n int
 fmt.Print("Введите номер числа в последовательности Фибоначчи: ")
 fmt.Scan(&n)

 result := fibonacci(n, memo)
 fmt.Printf("Число Фибоначчи под номером %d равно %d", n, result)
}
```

Задание 6. Обратные числа
   Напишите программу, которая принимает целое число и выводит его в обратном порядке. Например, для числа 12345 программа должна вывести 54321. Используйте цикл для обработки цифр числа.
```go
func main() {
 var number int
 fmt.Print("Введите целое число: ")
fmt.Scan(&number)

 reversedNumber := 0

 for number != 0 {
  remainder := number % 10
  reversedNumber = reversedNumber*10 + remainder
  number /= 10
  fmt.Println("Собираем обратное число:", reversedNumber)
 }
 fmt.Println("Обратное число:", reversedNumber)
}
```

Задание 7. Треугольник Паскаля
   Напишите программу, которая выводит треугольник Паскаля до заданного уровня. Для этого используйте цикл и массивы для хранения предыдущих значений строки треугольника.
```go
func main() {
 var level int

 fmt.Print("Введите уровень треугольника Паскаля: ")
 fmt.Scan(&level)

 triangle := make([][]int, level)

 for i := 0; i < level; i++ {
  triangle[i] = make([]int, i+1)
  triangle[i][0] = 1 // Первая единица в каждой строке
  triangle[i][i] = 1 // Последняя единица в каждой строке

  for j := 1; j < i; j++ {
   triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
  }
 }

 for i := 0; i < level; i++ {
 for k := level - i; k > 0; k-- {
   fmt.Print(" ")
  }
  for j := 0; j <= i; j++ {
   fmt.Print(triangle[i][j], " ")
  }
  fmt.Println()
 }
}

```

Задание 8. Число палиндром
   Напишите программу, которая проверяет, является ли число палиндромом (одинаково читается слева направо и справа налево). Не используйте строки для решения этой задачи — работайте только с числами.
```go
func isPalindrome(num int) bool {
 if num < 0 {
  return false
 }

 reversed := 0
 originalNum := num

 for num > 0 {
  remainder := num % 10
  reversed = reversed*10 + remainder
  num /= 10
 }

 return originalNum == reversed
}

func main() {
 var num int
 fmt.Println("Введите число:")
 fmt.Scan(&num)

 if isPalindrome(num) {
  fmt.Printf("%d является палиндромом", num)
 } else {
  fmt.Printf("%d не является палиндромом", num)
 }
}
```

Задание 9. Нахождение максимума и минимума в массиве
   Напишите функцию, которая принимает массив целых чисел и возвращает одновременно максимальный и минимальный элемент с использованием одного прохода по массиву.
```go
func MinMax(arr []int) (min int, max int) {    
 min = arr[0]
 max = arr[0]

 for _, value := range arr { // индекс нам не нужен
  if value < min {
   min = value
  }
  if value > max {
   max = value
  }
 }

 return min, max
}

func main() {
 arr := []int{3, 5, 1, 8, -2, 7}
 min, max := MinMax(arr)
 fmt.Printf("Минимальный элемент: %d, Максимальный элемент: %d", min, max)
}
```

Задание 10. Игра "Угадай число"
   Напишите программу, которая загадывает случайное число от 1 до 100, а пользователь пытается его угадать. Программа должна давать подсказки "больше" или "меньше" после каждой попытки. Реализуйте ограничение на количество попыток.
```go
func main() {
 rand.Seed(time.Now().UnixNano())
 randomNumber := rand.Intn(100) + 1 
 attempts := 10

 fmt.Println("Я загадал число от 1 до 100.")
 fmt.Println("У вас есть 10 попыток, чтобы его угадать.")

 for attempts > 0 {
  var userGuess int

  fmt.Print("Введите ваше предположение: ")
  fmt.Scan(&userGuess)

  if userGuess < 1 || userGuess > 100 {
   fmt.Println("Пожалуйста, введите число от 1 до 100.")
   continue
  }

  if userGuess < randomNumber {
   fmt.Println("Слишком мало! Попробуйте еще раз.")
  } else if userGuess > randomNumber {
   fmt.Println("Слишком много! Попробуйте еще раз.")
  } else {
   fmt.Printf("Поздравляю! Вы угадали загаданное число %d", randomNumber)
   return
  }

  attempts--
  fmt.Printf("У вас осталось %d попыток \n", attempts)

  if attempts == 0 {
   fmt.Printf("К сожалению, у вас закончились попытки. Загаданное число было %d", randomNumber)
  }
 }
}
```

Задание 11. Числа Армстронга
   Напишите программу, которая проверяет, является ли число числом Армстронга (число равно сумме своих цифр, возведённых в степень, равную количеству цифр числа). Например, 153 = 1³ + 5³ + 3³.
```go
func isArmstrong(num int) bool {
 strNum := fmt.Sprintf("%d", num)
 numDigits := len(strNum)

 sum := 0
 for _, digit := range strNum {
  d := int(digit - '0')
  sum += int(math.Pow(float64(d), float64(numDigits)))
 }

 return sum == num
}

func main() {
 var number int
 fmt.Print("Введите число: ")
 fmt.Scan(&number)

 if isArmstrong(number) {
  fmt.Printf("%d является числом Армстронга \n", number)
 } else {
  fmt.Printf("%d не является числом Армстронга \n", number)
 }
}
```

Задание 12. Подсчет слов в строке
   Напишите программу, которая принимает строку и выводит количество уникальных слов в ней. Используйте `map` для хранения слов и их количества.
```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Введите строку:")
  if scanner.Scan() {
    input := scanner.Text()
    words := strings.Fields(input)
    wordCount := make(map[string]int)
      for _, word := range words {
        wordCount[word]++
      }
    fmt.Printf("Количество уникальных слов: %d", len(wordCount))
    }
  else {
    fmt.Println("Ошибка считывания ввода.")
  }
}
```

Задание 13. Игра "Жизнь" (Conway's Game of Life)
   Реализуйте клеточный автомат "Жизнь" Конвея для двухмерного массива. Каждая клетка может быть либо живой, либо мертвой. На каждом шаге состояния клеток изменяются по следующим правилам:
   - Живая клетка с двумя или тремя живыми соседями остаётся живой, иначе умирает.
   - Мёртвая клетка с тремя живыми соседями оживает.
   Используйте циклы для обработки клеток.
```go
const (
    width  = 20 // ширина поля
    height = 10 // высота поля
)

type CellState bool

const (
    Dead  CellState = false // мёртвая клетка
    Alive CellState = true  // живая клетка
)

type Grid [height][width]CellState

// Метод для вычисления следующего поколения
func (g *Grid) NextGeneration() Grid {
    var newGrid Grid
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            liveNeighbours := g.countLiveNeighbours(x, y)
            if g[y][x] == Alive {
                // Живая клетка остаётся живой, если у неё 2 или 3 живых соседа
                if liveNeighbours == 2 || liveNeighbours == 3 {
                    newGrid[y][x] = Alive
                } else {
                    newGrid[y][x] = Dead
                }
            } else {
                // Мёртвая клетка становится живой, если у неё ровно 3 живых соседа
                if liveNeighbours == 3 {
                    newGrid[y][x] = Alive
                } else {
                    newGrid[y][x] = Dead
                }
            }
        }
    }
    return newGrid
}

// Метод для подсчёта живых соседей клетки
func (g *Grid) countLiveNeighbours(x, y int) int {
    count := 0
    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            if dx == 0 && dy == 0 {
                continue // пропускаем саму клетку
            }
            nx, ny := x+dx, y+dy
            if nx >= 0 && nx < width && ny >= 0 && ny < height {
                if g[ny][nx] == Alive {
                    count++
                }
            }
        }
    }
    return count
}

// Метод для вывода текущего состояния поля
func (g *Grid) Print() {
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if g[y][x] == Alive {
                fmt.Print("O")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func main() {
    var grid Grid

    // Инициализация начальной конфигурации (например, "глайдер")
    grid[1][2] = Alive
    grid[2][3] = Alive
    grid[3][1] = Alive
    grid[3][2] = Alive
    grid[3][3] = Alive

    for {
        fmt.Print("\033[H\033[2J") // Очистка экрана
        grid.Print()
        grid = grid.NextGeneration()
        time.Sleep(500 * time.Millisecond) // Пауза между поколениями
    }
}
```

Задание 14. Цифровой корень числа
   Напишите программу, которая вычисляет цифровой корень числа. Цифровой корень — это рекурсивная сумма цифр числа, пока не останется только одна цифра. Например, цифровой корень числа 9875 равен 2, потому что 9+8+7+5=29 → 2+9=11 → 1+1=2.
```go
func digitalRoot(n int) int {
 fmt.Printf("Промежуточное число %d \n", n)
 if n < 10 {
  return n
 }
 sum := 0
 for n > 0 {
  sum += n % 10
  n /= 10
 }
 return digitalRoot(sum)
}

func main() {
 var number int
 fmt.Print("Введите число: ")
 fmt.Scan(&number)
 result := digitalRoot(number)
 fmt.Printf("Цифровой корень числа %d равен %d", number, result)
}
```

Задание 15. Римские цифры
   Напишите функцию, которая преобразует арабское число (например, 1994) в римское (например, "MCMXCIV"). Программа должна использовать циклы и условные операторы для создания римской записи.
```go
func arabicToRoman(num int) string {
 val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
 sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

 roman := ""
 for i := 0; num > 0; i++ {
  for num >= val[i] {
   roman += sym[i] 
   num -= val[i]
  }
 }
 return roman
}

func main() {
 number := 1994
 romanNumeral := arabicToRoman(number)
 fmt.Printf("Арабское число %d в римских цифрах: %s", number, romanNumeral)
}
```

package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(task16("3BC9", 16, 2))
	fmt.Println(task17(2, 3, 1.125))
	fmt.Println(task18([]int{-5, 1, -10, 3, 2, -18}))
	fmt.Println(task19([]int{1, 5, 8, 19}, []int{0, 7, 9, 18, 21}))
	fmt.Println(task20("sub", "This is sub test string with substring"))
	fmt.Println(task21("9.111+2.2"))
	fmt.Println(task22("In girum imus, nocte et. consumimur igni"))
	fmt.Println(task23(1, 5, 3, 7, 2, 6))
	fmt.Println(task24("In girum imus, nocte et. consumimur igni"))
	fmt.Println(task25(2024))
	fmt.Println(task26(89))
	fmt.Println(task27(51, 100))
	fmt.Println(task28(300, 380))
	fmt.Println(task29("Hello world"))
	fmt.Println(task30(645, 381))
}

/*
Перевод чисел из одной системы счисления в другую
Напишите программу, которая принимает на вход число в произвольной системе счисления (от 2 до 36)
и переводит его в другую систему счисления.
*/
func task16(value string, inRadix int, outRadix int) string {
	if inRadix < 2 || inRadix > 36 || outRadix < 2 || outRadix > 36 {
		panic("Wrong radix!")
	}
	inter, err := strconv.ParseInt(value, inRadix, 64)
	if err != nil {
		panic(err)
	}
	return strconv.FormatInt(inter, outRadix)
}

/*
Решение квадратного уравнения
Реализуйте функцию для нахождения корней квадратного уравнения вида \( ax^2 + bx + c = 0 \).
Учтите случай комплексных корней.
*/
func task17(a float64, b float64, c float64) string {
	var D = math.Pow(b, 2) - 4*a*c
	var x1 = ""
	var x2 = ""
	if D > 0 {
		fmt.Println(D)
		x1 = strconv.FormatFloat((-b+math.Sqrt(D))/(2*a), 'f', -1, 64)
		x2 = strconv.FormatFloat((-b-math.Sqrt(D))/(2*a), 'f', -1, 64)
	} else if D == 0 {
		x1 = strconv.FormatFloat(-b/(2*a), 'f', -1, 64)
		x2 = x1
	} else {
		x1 = strconv.FormatFloat(-b/(2*a), 'f', -1, 64) + "+" + strconv.FormatFloat(
			math.Sqrt(-D)/(2*a), 'f', -1, 64) + "i"
		x2 = strconv.FormatFloat(-b/(2*a), 'f', -1, 64) + "-" + strconv.FormatFloat(
			math.Sqrt(-D)/(2*a), 'f', -1, 64) + "i"
	}
	return "x1 = " + x1 + " x2 = " + x2
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

/*
Сортировка чисел по модулю
Дан массив чисел. Напишите функцию, которая отсортирует его элементы по возрастанию их абсолютных значений.
*/
func task18(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		var minIdx = i
		for j := i; j < len(arr); j++ {
			if abs(arr[j]) < abs(arr[minIdx]) {
				minIdx = j
			}
		}
		if abs(arr[minIdx]) < abs(arr[i]) {
			var tmp = arr[i]
			arr[i] = arr[minIdx]
			arr[minIdx] = tmp
		}
	}
	return arr
}

/*
Слияние двух отсортированных массивов
Напишите функцию, которая объединяет два отсортированных массива в один, сохраняя их упорядоченность.
*/
func task19(arr1 []int, arr2 []int) []int {
	var out = make([]int, len(arr1)+len(arr2))
	var idx1 = 0
	var idx2 = 0
	for i := 0; i < len(arr1)+len(arr2); i++ {
		if (len(arr1)-1 < idx1) || ((len(arr2)-1 > idx2) && arr2[idx2] < arr1[idx1]) {
			out[i] = arr2[idx2]
			idx2++
		} else {
			out[i] = arr1[idx1]
			idx1++
		}
	}
	return out
}

/*
Нахождение подстроки в строке без использования встроенных функций
Реализуйте функцию, которая находит первую позицию вхождения одной строки в другую.
Если подстрока не найдена, вернуть -1.
*/
func task20(sub string, orig string) int {
	var idx = -1
	for i := 0; i < len(orig); i++ {
		if sub[0] != orig[i] {
			continue
		}
		if i+len(sub) > len(orig) {
			return -1
		}
		for j := 0; j < len(sub); j++ {
			if orig[i+j] != sub[j] {
				break
			} else if orig[i+j] == sub[j] && j == len(sub)-1 {
				idx = i
			}
		}
		if idx != -1 {
			return idx
		}
	}
	return idx
}

/*
Калькулятор с расширенными операциями
Напишите программу, которая выполняет различные математические операции (+, -, *, /, ^, %), заданные пользователем.
Реализуйте обработку ошибок, связанных с делением на ноль или недопустимой операцией.
*/
func task21(inp string) float64 {
	var num = regexp.MustCompile(`\d+\.\d+`)
	var name = regexp.MustCompile(`[+\-*/^%]`)
	var operator = name.FindAllString(inp, -1)
	if len(operator) != 1 {
		fmt.Println("Unknown operation!")
		return math.NaN()
	}
	var numbers = num.FindAllString(inp, -1)
	if len(numbers) != 2 {
		fmt.Println("Wrong input! Check your numbers!")
		return math.NaN()
	}
	var parsedNums = make([]float64, 2)
	for i := 0; i < len(numbers); i++ {
		parsedNums[i], _ = strconv.ParseFloat(numbers[i], 64)
	}
	var out = 0.0
	switch operator[0] {
	case "+":
		out = parsedNums[0] + parsedNums[1]
	case "-":
		out = parsedNums[0] - parsedNums[1]
	case "*":
		out = parsedNums[0] * parsedNums[1]
	case "/":
		out = parsedNums[0] / parsedNums[1]
	case "^":
		out = math.Pow(parsedNums[0], parsedNums[1])
	case "%":
		out = math.Mod(parsedNums[0], parsedNums[1])
	}
	return out
}

/*
Проверка палиндрома
Реализуйте функцию, которая проверяет, является ли строка палиндромом (игнорируя пробелы, знаки препинания и регистр)
*/
func task22(inp string) bool {
	var chars = regexp.MustCompile(`[^A-Za-z]`)
	inp = strings.ToLower(inp)
	var str = chars.ReplaceAllString(inp, "")
	var reverse = ""
	for _, v := range str {
		reverse = string(v) + reverse
	}
	return str == reverse
}

/*
Нахождение пересечения трех отрезков
Даны три отрезка на числовой оси (их начальные и конечные точки).
Нужно определить, существует ли область пересечения всех трех отрезков.
*/
func task23(x1 float64, x2 float64, x3 float64, x4 float64, x5 float64, x6 float64) bool {
	var maxStart = x1
	var minEnd = x2
	if x3 > maxStart {
		maxStart = x2
	}
	if x5 > maxStart {
		maxStart = x5
	}
	if x4 < minEnd {
		minEnd = x4
	}
	if x6 < minEnd {
		minEnd = x6
	}
	return maxStart <= minEnd
}

/*
Выбор самого длинного слова в предложении
Напишите программу, которая находит самое длинное слово в предложении, игнорируя знаки препинания.
*/
func task24(inp string) string {
	var chars = regexp.MustCompile(`[A-Za-z]+`)
	inp = strings.ToLower(inp)
	var str = chars.FindAllString(inp, -1)
	var maxIdx = 0
	for i := 0; i < len(str); i++ {
		if len(str[i]) > len(str[maxIdx]) {
			maxIdx = i
		}
	}
	return str[maxIdx]
}

/*
Проверка высокосного года
Реализуйте функцию, которая проверяет, является ли введенный год високосным по правилам григорианского календаря.
*/
func task25(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

/*
Числа Фибоначчи до определенного числа
Напишите программу, которая выводит все числа Фибоначчи, не превышающие заданного значения.
*/
func task26(n int) []int {
	if n < 0 {
		panic(errors.New("wrong input"))
	}
	var result = []int{0}
	if n == 0 {
		return result
	}
	result = append(result, 1)
	var idx = 2
	for {
		var next = result[idx-2] + result[idx-1]
		if next > n {
			break
		}
		result = append(result, next)
		idx++
	}
	return result
}

/*
Определение простых чисел в диапазоне
Реализуйте функцию, которая принимает два числа и выводит все простые числа в диапазоне между ними.
*/
func task27(start int, end int) []int {
	var numbers = make([]int, 0)
	var isPrime = false
	for i := start; i < end; i++ {
		isPrime = true
		for j := 2; j < int(math.Ceil(math.Sqrt(float64(i)))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

/*
Числа Армстронга в заданном диапазоне
Напишите программу, которая выводит все числа Армстронга в заданном диапазоне.
Число Армстронга – это число, равное сумме своих цифр, возведенных в степень, равную количеству цифр числа.
*/
func task28(start int, end int) []int {
	var numbers = make([]int, 0)
	for i := start; i < end; i++ {
		x := i
		n := 0
		for x != 0 {
			x /= 10
			n++
		}
		sum := 0
		x = i
		for x != 0 {
			digit := x % 10
			sum += int(math.Pow(float64(digit), float64(n)))
			x /= 10
		}
		if sum == i {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

/*
Реверс строки
Напишите программу, которая принимает строку и возвращает ее в обратном порядке,
не используя встроенные функции для реверса строк.
*/
func task29(in string) string {
	var out = ""
	for i := 0; i < len(in); i++ {
		out += string(in[len(in)-1-i])
	}
	return out
}

/*
Нахождение наибольшего общего делителя (НОД)
Реализуйте алгоритм Евклида для нахождения наибольшего общего делителя двух чисел с использованием цикла.
*/
func task30(x int, y int) int {
	if x < y {
		for {
			z := y % x
			if z == 0 {
				break
			}
			y = x
			x = z
		}
		return x
	} else {
		for {
			z := x % y
			if z == 0 {
				break
			}
			x = y
			y = z
		}
		return y
	}
}
