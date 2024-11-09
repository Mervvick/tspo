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
