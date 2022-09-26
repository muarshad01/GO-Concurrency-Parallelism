/* heatEqn.go */
package main

import ( "fmt"; "os"; "runtime"; "strconv"; "sync" )

func newVal(M [][]float32, N [][]float32,
   from int, to int, n int, waitF *sync.WaitGroup) {
   for i := from; i < to; i++ {
      for j := 1; j < n-1; j++ {
         N[i][j] = (4.0*M[i+1][j] +
                    4.0*M[i-1][j] +
                    4.0*M[i][j+1] +
                    4.0*M[i][j-1] +
                    M[i+1][j+1]   +
                    M[i-1][j-1]   +
                    M[i+1][j-1]   +
                    M[i-1][j+1]) / 20.0
      }
   }
   waitF.Done()
}

func partition(n int, nf int) (int, int, int, int) {
   size1 := n / nf
   size2 := size1 + 1
   n2 := n % nf
   n1 := nf - n2
   valid := (size1*n1 + size2*n2) == n
   if !valid {
      fmt.Println("partition error\n")
      os.Exit(1)
   }
   return size1, n1, size2, n2
}

func dump(M [][]float32, n int) {
   for j := n-1; j >= 0; j-- {
      for i := 0; i < n; i++ {
         fmt.Printf("%7.2f", M[i][j])
      }
      fmt.Printf("\n")
   }
   fmt.Printf("\n")
}

func main() {
   if len(os.Args) < 5 {
      fmt.Printf(
         "<filename> n,numGoroutines,maxProcs,iters\n")
      return
   }
   n, _             := strconv.Atoi(os.Args[1])
   numGoroutines, _ := strconv.Atoi(os.Args[2])
   maxProcs, _      := strconv.Atoi(os.Args[3])
   iters, _         := strconv.Atoi(os.Args[4])
   runtime.GOMAXPROCS(maxProcs)

   M := make([][]float32, n)
   N := make([][]float32, n)
   for i := 0; i < n; i++ {
      M[i] = make([]float32, n)
      N[i] = make([]float32, n)
   }
   for i := 0; i < n; i++ {
      M[0][i] = 100.0
      M[i][n-1] = 100.0
      N[0][i] = 100.0
      N[i][n-1] = 100.0
   }
   var waitF sync.WaitGroup
   //start := time.Now()
   size1, n1, size2, _ := partition(n-2, numGoroutines)
   for i := 0; i < iters; i++ {
      from := 1
      to := 0
      waitF.Add(numGoroutines)
      for f := 0; f < numGoroutines; f++ {
         if f < n1 {
            to = from + size1
         } else {
            to = from + size2
         }
         if i%2 == 0 {
            go newVal(M, N, from, to, n, &waitF)
         } else {
            go newVal(N, M, from, to, n, &waitF)
         }
         from = to
      }
      waitF.Wait()
   }
   //elapse := time.Since(start)
   dump(M, n)
   //fmt.Printf("n %6d sec %7.4f nsec/pt %6.4f\n",
   //n,
   //float32(elapse.Seconds()),
   //float32(elapse.Nanoseconds())/float32(iters*n*n))
}
