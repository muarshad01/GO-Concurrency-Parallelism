/* function.go */
package main
import "fmt"

func myFunc(i int){
   fmt.Printf("hello from function %2d\n", i);
}

func main() {
   for i := 0; i < 10; i++ {
      myFunc(i)
   }
}
