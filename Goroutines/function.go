/* function.go: serial functions*/

package main

import ("fmt"; "math/rand"; "os"; "strconv")

func myFunction(id int) {
   
   fmt.Printf("\nmyFunction %2d:",id)
   
   var sum float64
   for i := 0; i < rand.Intn(100000); i++ {
      sum += rand.Float64()
   }
   fmt.Printf("\ndone %7.3e id %2d", sum, id)
}

func main() {

   if len(os.Args) < 2 {
      fmt.Printf("usage <filename> seed\n")
      return
   }

   seed, _  := strconv.Atoi(os.Args[1])
   rand.Seed(int64(seed))

   fmt.Printf("\n--start main program--")
   for id := 0; id < 20; id++ {
      myFunction(id)
   }
   fmt.Printf("\n--end main program--\n")
}
