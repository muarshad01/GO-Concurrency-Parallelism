/* anonFunc.go */
package main
import "fmt"

func main() {
   for i := 0; i < 10; i++ {
      func(i int){
         fmt.Printf("hello from function %2d\n", i);
      }(i)
   }
}
