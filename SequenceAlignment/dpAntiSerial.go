/* dpAntiSerial.go */
package main

import ( "fmt"; "log"; "math/rand"; "os"; "runtime"; "strconv"; "time")

func MIN(x, y int32) int32 {
   if x > y {
      return y
   }
   return x
}

func unEqual(P, T byte) int32 {
   if P != T {
      return 1
   }
   return 0
}

func dump(D [][]int32, P []byte, T []byte, n int32) {
   var i, j int32
   fmt.Printf("   ")
   for i = 0; i <= n; i++ {
      fmt.Printf("%3c", T[i])
   }
   fmt.Printf("\n")
   for i = 0; i <= n; i++ {
      fmt.Printf("%3c", P[i])
      for j = 0; j <= n; j++ {
         fmt.Printf("%3d", D[i][j])
      }
      fmt.Printf("\n")
   }
   fmt.Printf("\n")
}

func align(D [][]int32, P []byte, T []byte, n int32) {
   aP := make([]byte, 2*n)
   aT := make([]byte, 2*n)
   i := n;  j := n; Tp := n; Pp := n
   pPtr := 0; tPtr := 0; var length int32 = 0
   fmt.Printf("(%d,%d)",i,j)
   for {
      if(i>0)&&(D[i][j]==D[i-1][j]+1){
         aT[tPtr] = '-'; if(i>0){i--}; tPtr++
         aP[pPtr] = P[Pp]; if(Pp>0){Pp--}; pPtr++
      } else {
         if(j>0)&&(D[i][j]==D[i][j-1]+1){
            aP[pPtr] = '-'; if(j>0){j--}; pPtr++
            aT[tPtr] = T[Tp]; if (Tp>0){Tp--}; tPtr++
         } else {
            aP[pPtr] = P[Pp]; if(i>0){i--}; if(Pp>0){Pp--}; pPtr++
            aT[tPtr] = T[Tp]; if(j>0){j--}; if(Tp>0){Tp--}; tPtr++
         }
      }
      length++
      fmt.Printf("(%d,%d)",i,j)
      if((i==0)&&(j==0)){
         fmt.Printf("\nn = %2d aligned length = %2d distance = %2d\n", n, length, D[n][n]);
         break
      }
   }

   fmt.Print("  ")
   for i=0; i<length; i++{
      fmt.Printf("%3d", i)
   }
   fmt.Printf("\n  ")
   for i=1; i<=length; i++{
      fmt.Printf("%3c", aP[length-i])
   }
   fmt.Printf("\n  ")
   for i=1; i<=length; i++{
      fmt.Printf("%3c", aT[length-i])
   }
   fmt.Printf("\n")

}

func main() {
   if len(os.Args) < 4 {
      fmt.Printf("need:  n seed maxGB\n")
      return
   }
	n64,_    := strconv.ParseInt(os.Args[1],10,64)
   n        := int32(n64)
   seed,_   := strconv.ParseInt(os.Args[2],10,64)
   kBytes,_ := strconv.ParseInt(os.Args[3],10,64)
   maxMem  := uint64(float32(kBytes)*0.75*1024.0)
   rand.Seed(seed)

   start := time.Now()
   var i, j int32
   P := make([]byte,n+1); T := make([]byte,n+1)
   P[0] = ' ';            T[0] = ' '
   dna := [4]byte{'a','c','t','g'}
   for i = 1; i < n+1; i++ {
      P[i] = byte(dna[rand.Intn(4)])
      T[i] = byte(dna[rand.Intn(4)])
   }
   var mem runtime.MemStats
   D := make([][]int32, n+1)
   for i = 0; i < n+1; i++ {
      D[i] = make([]int32, n+1)
      runtime.ReadMemStats(&mem)
      if(uint64(mem.Sys) > maxMem){
         log.Printf("%s abort: memory limit exceeded %v>%v\n",os.Args[0],mem.Sys, maxMem)
         os.Exit(1)
      }
      D[i][0] = i; D[0][i] = i
   }
   for j=1; j<=n+1; j++ {
      for i=1; i<j; i++ {
         D[i][j-i]=MIN(MIN(D[i-1][j-i]+1, D[i][j-i-1]+1), D[i-1][j-i-1]+ unEqual(P[i],T[j-i]))
      }
   }
   for i=1; i<=n; i++ {
      for j=n-i; j>0; j-- {
         D[n-j+1][j+i]=MIN(MIN(D[n-j][j+i]+1, D[n-j+1][j+i-1]+1), D[n-j][j+i-1]+ unEqual(P[n-j+1],T[j+i]))
      }
   }

   elapse := time.Since(start)
   if(n<=20){
      dump(D, P, T, n)
      align(D, P, T, n)
   }
   fmt.Println(n, elapse.Seconds())
   os.Exit(0)
}
