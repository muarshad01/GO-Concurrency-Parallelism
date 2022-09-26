/*channelEx1.go*/
package main

import ("fmt"; "os"; "strconv")

var done = make(chan bool)

func routine1(p chan int, n int) {
   for i:=0; i<n; i++{
      p <- i
      fmt.Println("routine 1 sent", i)
      s := <-p
      fmt.Println("routine 1 received", s)
   }
   close(p)
   fmt.Println("routine 1 ended")
}

func routine2(p chan int) {
   for {
      v, open := <-p
      if !open {
         fmt.Println("routine2 ended")
         done <- true
         return
      }
      fmt.Println("routine 2 received", v)
      p <- v*v
      fmt.Println("routine 2 sent", v*v)
   }
}

func main() {
   n, _ := strconv.Atoi(os.Args[1])
   p := make(chan int)
   go routine1(p, n)
   go routine2(p)
   <-done
   fmt.Println("main program ended")
}
