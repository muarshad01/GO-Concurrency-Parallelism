/* serialSieve.go */
package main
import ( "fmt"; "os"; "strconv"; "time")
func main() {
   N, _  := strconv.Atoi(os.Args[1])
   start := time.Now()
   Prime := make([]bool, N+1)
   for i := range Prime {
      Prime[i] = true
   }
   for i := 2; i*i <= N; i++ {
      if Prime[i] {
         for j := i * 2; j <= N; j = j + i {
            Prime[j] = false
         }
      }
   }
   elapse := time.Since(start)
   count := 0
   for i := 2; i <= N; i++ {
      if Prime[i] {
         fmt.Printf("%d\n", i)
         count++
      }
   }
   fmt.Println(count, elapse.Microseconds())
}
