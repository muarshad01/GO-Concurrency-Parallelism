/* program3.go */

package main

import ( "fmt"; "os"; "strconv" )

func main() {

   if len(os.Args) != 3 {
      fmt.Println("Two arguments expected")
      return
   }

   m, _ := strconv.Atoi(os.Args[1])
   n, _ := strconv.Atoi(os.Args[2])

   for i:=0; i<m; i++ {
      for j:=0; j<n; j++ {
         fmt.Printf("%3d", i+j)
      }
      fmt.Printf("\n");
   }
}

