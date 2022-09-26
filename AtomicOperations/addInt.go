/* addInt.go */
package main

import ("fmt"; "os"; "runtime"; "strconv"; "sync")

func addToSum(theSum *int64, val int64, n int, wg *sync.WaitGroup) {

   for i := 0; i < n; i++ {
      *theSum += val
   }
   wg.Done()
}

func main() {

   n, _        := strconv.Atoi(os.Args[1])
   numFuncs, _ := strconv.Atoi(os.Args[2])
   Procs, _    := strconv.Atoi(os.Args[3])

   var theSum int64 = 0
   var wg sync.WaitGroup
   runtime.GOMAXPROCS(Procs)

   wg.Add(int(numFuncs))
   for f := 0; f < numFuncs; f++ {
      go addToSum(&theSum, int64(1), n, &wg)
   }
   wg.Wait()

   fmt.Println("n", n, "Funcs", numFuncs, "Procs", Procs, "sum", theSum)
}
