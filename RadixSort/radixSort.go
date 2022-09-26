/* radixSort.go */
package main

import ( "fmt"; "math/rand"; "os"; "strconv"; "sort" )

func getMax(arr[]int) int {
   mx := arr[0]
   for i := 1; i < len(arr); i++ {
      if (arr[i] > mx){
         mx = arr[i]
      }
   }
   return mx
}

func countingSort(arr[]int, exp int ) {
   n      := len(arr)
   output := make([]int,n)   // output array
   count  := make([]int,11)  // initial 0 plus 0-9 digits

   for i := 0; i < n; i++{   // Store occurrences in count[] 
      count[(arr[i]/exp)%10]++
   }

   fmt.Printf("count  %3d\n",count)

   sum := 0
   for i := 0; i < 11; i++ { // Exclusive prefix sums
      t := count[i]
      count[i] = sum
      sum += t
   }
   fmt.Printf("scan   %3d\n",count);

   for i := 0; i < n; i++ {  //copy to output
      output[count[(arr[i]/exp)%10]] = arr[i]
      count[(arr[i]/exp)%10]++
   }
   fmt.Printf("output %3d\n",output);

   for i := 0; i < n; i++{     //copy back 
      arr[i] = output[i]
   }
}

func radixsort(arr[]int) {
   m := getMax(arr)                            //max number: we need to know digits
   for exp := int(1); m/exp > 0; exp *= 10 {   //counting sort for each digit
      countingSort(arr, exp);                  //exp is 10^i (i is digit number)
   }
}

func main() {
   n, _       := strconv.Atoi(os.Args[1])
   seed, _    := strconv.Atoi(os.Args[2])
   numSize, _ := strconv.Atoi(os.Args[3])
   arr := make([]int, n)
   aux := make([]int, n)
   rand.Seed(int64(seed))

   for i:=0; i<n; i++{
      arr[i] = rand.Intn(numSize)
      aux[i] = arr[i];   //to test sorting
   }

   fmt.Printf("array  %3d\n",arr)
   radixsort(arr)
   fmt.Printf("sorted %3d\n",arr)
   // test if arr was properly sorted, using GO's own sort
   sort.Ints(aux)
   for i:=0; i<n; i++ {
      if(arr[i]!=aux[i]){
         fmt.Println("failed")
         return
      }
   }
   fmt.Println("ok")
}
