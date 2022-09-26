/* goroutine2.go
   concurrent functions with synchronization */

package main

import ("fmt"; "sync"; "math/rand"; "os"; "strconv")

func myFunction(id int, numFuncs *sync.WaitGroup) {

   fmt.Printf("\nmyFunction %2d:",id)
   var sum float64
   for i := 0; i < rand.Intn(100000); i++ {
      sum+=rand.Float64()
   }
   fmt.Printf("\ndone %7.3e id %2d",sum,id)
   numFuncs.Done()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("usage <filename> seed\n")
		return
	}

   seed, _ := strconv.Atoi(os.Args[1])
   rand.Seed(int64(seed))

   var numFuncs sync.WaitGroup
   fmt.Printf("\n--start main program--")
   numFuncs.Add(20)
   for i := 0; i < 20; i++ {
      go myFunction(i, &numFuncs)
   }
   numFuncs.Wait()
   fmt.Printf("\n--end main program--\n")
}
