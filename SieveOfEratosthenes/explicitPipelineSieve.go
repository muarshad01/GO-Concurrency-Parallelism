/*explicitPipelineSieve.go*/
package main
import ( "fmt"; "math"; "os"; "strconv"; "time")

var done = make(chan bool)

func source(p chan int, N int) {
   for i := 2; i <= N; i++ {
      p <- i
   }
   close(p)
}

func filter(index int, in chan int, out chan int) {
   var ( counter = 0; prime   int)
   for {
      val, open := <-in
      if !open {
         close(out)
         return
      }
      switch {
      case (counter <= index):
         out <- val
         if counter == index {
            prime = val
         }
      default:
         if val%prime != 0 {
            out <- val
         }
      }
      counter++
   }
}

func sink(p chan int, count *int) {
   for {
      val, open := <-p
      if !open {
         done <- true
         return
      } else {
         //fmt.Println(val)
         *count++
      }
   }
}

func main() {
   N, _ := strconv.Atoi(os.Args[1])
   numChannels:= int((math.Sqrt(float64(N))))
   var channel= make([]chan int, numChannels)
   start := time.Now()
   for p := 0; p < numChannels; p++ {
      channel[p] = make(chan int)
   }
   go source(channel[0], N)
   for i := 0; i < numChannels-1; i++ {
      go filter(i, channel[i], channel[i+1])
   }
   var count = 0
   go sink(channel[numChannels-1], &count)
   <-done
   elapse := time.Since(start)
   fmt.Println(count, elapse.Microseconds())
}
