/* program4.go */

package main

import ( "fmt"; "os"; "strconv" )

func main() {

   if len(os.Args) < 2 {
      fmt.Println("Two arguments expected")
      return
   }

   m, _ := strconv.Atoi(os.Args[1])
   n, _ := strconv.Atoi(os.Args[2])

   for i:=0; i<m; i++ {
      for j:=0; j<n; j++ {
         fmt.Printf("*")
      }
      fmt.Printf("\n");
   }
   fmt.Printf("\n");
}

