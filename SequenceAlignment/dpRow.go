/* dpRow.go */
package main

import "fmt"

func MIN(x, y int) int {
   if x > y {
      return y
   }
   return x
}

func unEqual(P, T byte) int {
   if P != T {
      return 1
   }
   return 0
}

func dump(D [][]int, P []byte, T []byte, n int) {
   fmt.Printf("   ")
   for i := 0; i <= n; i++ {
      fmt.Printf("%3c", T[i])
   }
   fmt.Printf("\n")
   for i := 0; i <= n; i++ {
      fmt.Printf("%3c", P[i])
      for j := 0; j <= n; j++ {
         fmt.Printf("%3d", D[i][j])
      }
      fmt.Printf("\n")
   }
   fmt.Printf("\n")
}

func align(D [][]int, P []byte, T []byte, n int) {
   aP := make([]byte, 2*n)
   aT := make([]byte, 2*n)
   i := n;  j := n; Tp := n; Pp := n
   pPtr := 0; tPtr := 0; length := 0
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

   var SP, ST string
   fmt.Scanln(&SP)
   fmt.Scanln(&ST)
   if len(SP) != len(ST) {
      fmt.Println("strings have unequal length\n")
      return
   }
   n := len(SP)
   D := make([][]int, n+1)
   for i := 0; i < n+1; i++ {
      D[i] = make([]int, n+1)
      D[i][0] = i
      D[0][i] = i
   }
   P := make([]byte, n+1)
   T := make([]byte, n+1)
   P[0] = ' '
   T[0] = ' '
   for i := 1; i < n+1; i++ {
      P[i] = SP[i-1]
      T[i] = ST[i-1]
   }

   for i := 1; i <= n; i++ {
      myPi := P[i]
      for j := 1; j <= n; j++ {
         D[i][j] = MIN(MIN(D[i-1][j]+1, D[i][j-1]+1), D[i-1][j-1]+unEqual(myPi, T[j]))
      }
   }

   dump(D, P, T, n)
   align(D, P, T, n)
}
