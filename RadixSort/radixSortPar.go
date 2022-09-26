/* radixSortPar.go */
package main

import ( "fmt"; "math"; "math/rand"; "os"; "runtime"; "strconv";
          "sync"; "sync/atomic"; "sort" )

func getMax(arr[]int64) int64 {
   theMax := arr[0]
   for i := 1; i < len(arr); i++{
      if (arr[i] > theMax){
         theMax = arr[i]
      }
   }
   return theMax
}

func Min(x, y int64) int64 {
   if x < y {
      return x
   }
   return y
}

func makeCount(arr[]int64, count[]int64, chunk int64, exp int64, f int64, wg *sync.WaitGroup){
   n:= int64(len(arr))
   for i := f*chunk; i < Min((f+1)*chunk,n); i++{
      atomic.AddInt64(&count[(arr[i]/exp)%10],int64(1))
   }
   wg.Done()
}

func copyBack(arr[]int64, output[]int64, chunk int64, f int64, wg *sync.WaitGroup){
   n:= int64(len(arr))
   for i := f*chunk; i < Min((f+1)*chunk,n); i++{
      arr[i] = output[i]
   }
   wg.Done()
}

func makeOutput(arr[]int64, output[]int64, count[]int64,
                chunk int64, exp int64, f int64, wg *sync.WaitGroup){
   n := int64(len(arr))
   for i := f*chunk; i < Min((f+1)*chunk,n); i++{
      loc := (arr[i]/exp)%(int64(10))
      for{
         val := count[loc]
         if(atomic.CompareAndSwapInt64(&count[loc], val, val+1)){
            output[val]=arr[i]
            break
         }
      }
   }
   wg.Done()
}

func countingSort(arr[]int64, aux[]int, exp int64, funcs int64, Procs int64) {
   n      := int64(len(arr))
   output := make([]int64,n)
   count  := make([]int64,11)
   chunk  := int64(math.Ceil(float64(n)/float64(funcs)))
   var wg sync.WaitGroup

   wg.Add(int(funcs))
   for f:=int64(0); f<funcs; f++{
      go makeCount(arr, count, chunk, exp, f, &wg)
   }
   wg.Wait()

   fmt.Printf("count  %3d\n",count)

   sum := int64(0)
   for i := 0; i < 11; i++{    // Exclusive prefix sums, serial
      t := count[i]
      count[i] = sum
      sum += t
   }
   fmt.Printf("scan   %3d\n",count)

   wg.Add(int(funcs))
   for f:=int64(0); f<funcs; f++{
      go makeOutput(arr, output, count, chunk, exp, f, &wg)
   }
   wg.Wait()
   fmt.Printf("output %3d\n\n",output)

   wg.Add(int(funcs))
   for f:=int64(0); f<funcs; f++{
      go copyBack(arr, output, chunk, f, &wg)
   }
   wg.Wait()
}

func radixsort(arr[]int64, aux[]int, funcs int64, Procs int64) {
   m := getMax(arr)      //max number: we need to know digits
   for exp := int64(1); m/exp > 0; exp *= 10 {     //counting sort for each digit
      countingSort(arr, aux, exp, funcs, Procs);   //exp is 10^i (i is digit number)
   }
}

func main() {
   n,_       := strconv.ParseInt(os.Args[1],10,64)
   seed,_    := strconv.ParseInt(os.Args[2],10,64)
   numSize,_ := strconv.ParseInt(os.Args[3],10,64)
   funcs,_   := strconv.ParseInt(os.Args[4],10,64)
   Procs,_   := strconv.ParseInt(os.Args[5],10,64)
   rand.Seed(seed)
   runtime.GOMAXPROCS(int(Procs))
   arr := make([]int64,n)
   aux := make([]int,n)
   for i:=int64(0); i<int64(n); i++{
      arr[i] = rand.Int63n(int64(numSize))
      aux[i] = int(arr[i])
   }

   fmt.Printf("array  %3d\n\n",arr)
   radixsort(arr, aux, funcs, Procs)
   sort.Ints(aux)
   result:="ok"
   for i:=int64(0);i<n;i++{
      if(arr[i]!=int64(aux[i])){
         result="failed"
      }
   }
   fmt.Printf("%v n %4d funcs %2d Procs %2d seed %2d numSize %3d\n",
               result,n, funcs, Procs, seed, numSize)
}
