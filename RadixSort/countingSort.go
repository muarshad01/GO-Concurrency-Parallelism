/* countingSort.go */
package main

import ( "fmt"; "math/rand"; "os"; "strconv"; "sort" )

func getMax(arr[]int) int {
   theMax := arr[0]
   for i := 1; i < len(arr); i++ {
      if arr[i] > theMax {
         theMax = arr[i]
      }
   }
   return theMax
}

func countingSort(arr[]int, output[]int, m int) {
   n := len(arr)
   count := make([]int, m)

   for i := 0; i < n; i++ { 			// 1. Store occurrences in count[]
      count[arr[i]]++
   }
   fmt.Printf("count  %3d\n",count);

   sum := 0
   for i := 0; i < m; i++ { 			// 2. Exclusive prefix sums
      t := count[i]
      count[i] = sum
      sum += t
   }
   fmt.Printf("scan   %3d\n",count);

   for i := 0; i < n; i++ { 			//3. copy to output
      output[count[arr[i]]] = arr[i]
      count[arr[i]]++
   }
}

func main() {
   n, _       := strconv.Atoi(os.Args[1])
   seed, _    := strconv.Atoi(os.Args[2])
   numSize, _ := strconv.Atoi(os.Args[3])
   arr    := make([]int, n)
   aux    := make([]int, n)
   output := make([]int, n)
   rand.Seed(int64(seed))

   for i:=0; i<n; i++{
      arr[i] = rand.Intn(numSize)
      aux[i] = arr[i]
   }

   fmt.Printf("arr    %3d\n",arr);
   theMax := getMax(arr)+2  //max number plus 2 extra spaces
   countingSort(arr, output, theMax)
   fmt.Printf("sorted %3d\n",output)
   // test if arr was properly sorted, using GO's own sort
   sort.Ints(aux)
   for i:=0; i<n; i++ {
      if output[i] != aux[i] {
         fmt.Println("failed")
         return
      }
   }
   fmt.Println("ok")
}
