/* loadArrayPar.go */
package main

import ("fmt"; "os"; "runtime"; "strconv"; "sync"; "sync/atomic")

func makeHist(arr[]int64){
   theHist := make([]int64,len(arr))
   for i:=0; i<len(arr);i++{
      theHist[arr[i]]++
   }
   for i:=0; i<len(arr);i++{
      if theHist[i]>1 {
         fmt.Printf("%7d: %3d\n",i,theHist[i])
      }
   }
   fmt.Println()
}

func main() {
   n, _      := strconv.ParseInt(os.Args[1],10,64)
   nFuncs, _ := strconv.ParseInt(os.Args[2],10,64)
   Procs, _  := strconv.Atoi(os.Args[3])
   runtime.GOMAXPROCS(Procs)

   theArr :=  make([]int64,n*nFuncs)
   theOut :=  make([]int64,n*nFuncs)
   var wg sync.WaitGroup

   for i:=int64(0);i<nFuncs*n;i++{
      theArr[i]=i
   }

   place:=int64(0)
   wg.Add(int(nFuncs))
   for f:=int64(0); f<nFuncs; f++ {
      go func(wg *sync.WaitGroup){
         for i := int64(0); i < n; i++{
            for{
               myplace:=place
               if(atomic.CompareAndSwapInt64(&place, myplace, myplace+1)){
                  theOut[myplace] = theArr[myplace]
                  break
               }
            }
         }
         wg.Done()
      }(&wg)
   }
   wg.Wait()

   fmt.Println("n", n, "nFuncs", nFuncs, "out", theOut, "len", len(theArr))
   makeHist(theOut)
}
