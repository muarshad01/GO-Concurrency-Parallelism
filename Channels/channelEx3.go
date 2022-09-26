/*channelEx2.go*/
package main

import ("fmt"; "os"; "strconv"; "math/rand"; "time")

func client(c chan int) {
   for {
      num := rand.Intn(10)
      c <- num
      fmt.Printf("client sent %d\n",  num)
      val :=  <- c
      fmt.Printf("client received %d\n",  val)
      time.Sleep(2*time.Second)
   }
}

func server(ch []chan int) {

   var n int

   for {
      select {
      case n = <- ch[0]:
         fmt.Printf("Server received %3d from client 0\n", n)
         ch[0] <- n*2
      case n = <- ch[1]:
         fmt.Printf("Server received %3d from client 1\n", n)
         ch[1] <- n*3
      case n = <- ch[2]:
         fmt.Printf("Server received %3d from client 2\n", n)
         ch[2] <- n*4
      }
   }
}

func main() {

	seed, _ := strconv.Atoi(os.Args[1])
   	rand.Seed(int64(seed))
	
	var ch = make([]chan int, 3)
	for p := 0; p < 3; p++ {
		ch[p] = make(chan int)
	}
	
	go server(ch)

	for c := 0; c < 3; c++ {
	   go client(ch[c]);
	}

   	fmt.Scanln()
	fmt.Println("main program ended")
}
