package main

import "fmt"

func main() {
  fmt.Println("Hello, Go!")
}
// TASK 1
package main

import "fmt"

func main() {
    fmt.Println("I like Go!")
}
// TASK 2

package main

import "fmt"
 
func main() {
  var name string
  var age int
  fmt.Println("Введите имя")
  fmt.Scan(&name)
  fmt.Println("Введите возраст")
  fmt.Scan(&age)

  fmt.Println(name, age)
     

}

package main
import "fmt"
func main(){

  var a int
  var b int
  var a_1 int
  var b_1 int
  var c int
  fmt.Scan(&a) // считаем переменную 'a' с консоли
  fmt.Scan(&b) // считаем переменную 'b' с консоли
  a_1 = a * a
  b_1 = b * b
  c = a_1 + b_1
  fmt.Println(c)
}
