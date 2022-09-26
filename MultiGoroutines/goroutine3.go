/* goroutine3.go
   concurrent functions with even-odd sync */

package main

import ("fmt"; "sync"; "math/rand"; "os"; "strconv")

func myFunction(id int, numEvenFuncs, numOddFuncs *sync.WaitGroup) {
   fmt.Printf("\nmyFunction %2d:", id)
   var sum float64
   for i := 0; i < rand.Intn(10000); i++ {
      sum += rand.Float64()
   }

   fmt.Printf("\ndone %7.3e id %2d", sum, id)

   switch id%2 {
      case 0: numEvenFuncs.Done()
      case 1: numOddFuncs.Done()
   }
}

func main() {

   if len(os.Args) < 2 {
      fmt.Printf("usage <filename> seed\n")
      return
   }

   seed, _ := strconv.Atoi(os.Args[1])
   rand.Seed(int64(seed))

   var numEvenFuncs, numOddFuncs sync.WaitGroup

   fmt.Printf("\n--start main program--")
   numEvenFuncs.Add(10)
   numOddFuncs.Add(10)
   for i := 0; i < 20; i++ {
      go myFunction(i, &numEvenFuncs, &numOddFuncs)
   }
   numEvenFuncs.Wait()
   fmt.Printf("\n--end EVEN goroutines--\n")
   numOddFuncs.Wait()
   fmt.Printf("\n--end ODD goroutines--\n")
   fmt.Printf("\n--end main program--\n")
}
