/*channelEx2.go*/
package main

import ("fmt"; "os"; "strconv"; "math/rand"; "time")

func client(c chan int, id int) {
  for {
    num := rand.Intn(10)
    c <- num
    fmt.Printf(
       "client %d sent %d\n", id, num)
    val :=  <-c
    fmt.Printf(
       "client %d rcvd %d\n", id, val)
    time.Sleep(2*time.Second)
  }
}

func server(ch0 chan int,
            ch1 chan int,
	    ch2 chan int){
  var n int
  for {
    select {
    case n = <-ch0:
      fmt.Printf(
        "Server rcvd %3d from client 0\n", n)
      ch0 <- n*2
    case n = <-ch1:
      fmt.Printf(
        "Server rcvd %3d from client 1\n", n)
      ch1 <- n*3
    case n = <-ch2:
      fmt.Printf(
        "Server rcvd %3d from client 2\n", n)
      ch2 <- n*4
    }
  }
}

func main() {

  seed, _ := strconv.Atoi(os.Args[1])
  rand.Seed(int64(seed))

  ch0 := make(chan int)
  ch1 := make(chan int)
  ch2 := make(chan int)

  go server(ch0, ch1, ch2)
  
  go client(ch0, 0)
  go client(ch1, 1)
  go client(ch2, 2)

  fmt.Scanln()
  fmt.Println("main program ended")
}
