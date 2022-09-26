/* logSensors.go */
package main

import ("fmt"; "math/rand"; "os"; "runtime"; "strconv"; "sync"; "sync/atomic")

func makeHist(arr[]int64, totalObs int64, nFuncs int64){
   theHist := make([]int64,len(arr))
   for i:=0; i<len(arr);i++{
      theHist[arr[i]]++
   }
   for i:= int64(0);i < nFuncs; i++{
      fmt.Printf("%4d: %3d,",i,theHist[i])
   }
   for i:=0; i<len(arr);i++{
      totalObs -= theHist[i]
   }
   if totalObs==0{
       println("\nObservations consistent")
   }
}

func main() {
   maxObserv, _ := strconv.ParseInt(os.Args[1],10,64)
   seed, _      := strconv.ParseInt(os.Args[2],10,64)
   nFuncs, _    := strconv.ParseInt(os.Args[3],10,64)
   Procs, _     := strconv.Atoi(os.Args[4])
   runtime.GOMAXPROCS(Procs)
   rand.Seed(seed)

   var wg sync.WaitGroup
   theLog   := make([]int64,nFuncs*maxObserv)
   place    := int64(0)
   totalObs := int64(0)

   wg.Add(int(nFuncs))
   for f:=int64(0); f<nFuncs; f++ {
      go func(f int64, wg *sync.WaitGroup){
         numObserv := rand.Int63n(maxObserv)
         for i := int64(0); i < numObserv; i++{
            for{
               myplace:=place
               if(atomic.CompareAndSwapInt64(&place, myplace, myplace+1)){
                  theLog[myplace] = f
                  break
               }
            }
         }
         totalObs=atomic.AddInt64(&totalObs, numObserv)
         wg.Done()
      }(f,&wg)
   }
   wg.Wait()

   fmt.Println(nFuncs, Procs, totalObs)
   fmt.Println("log",theLog[0:totalObs])
   makeHist(theLog[0:totalObs], totalObs, nFuncs)
}
