/* goroutine4.go
   synchronizing an arbitrary graph */

package main

import ("fmt"; "sync"; "math/rand"; "os"; "strconv"; "time")

func START() {
   fmt.Printf("\nSTART")
}

func STOP() {
   fmt.Printf("\nSTOP")
}

func alpha(wg *sync.WaitGroup) {
   
   time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
   
   fmt.Printf("\nalpha")
   wg.Done()
}

func beta(wg *sync.WaitGroup) {
   
   time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
   
   fmt.Printf("\nbeta")
   wg.Done()
}

func gamma(wg *sync.WaitGroup) {
   
   time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
   
   fmt.Printf("\ngamma")
   wg.Done()
}

func delta(wg *sync.WaitGroup) {
   
   time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
   
   fmt.Printf("\ndelta")
   wg.Done()
}

func epsilon(wg *sync.WaitGroup) {
   
   time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
   
   fmt.Printf("\nepsilon")
   wg.Done()
}

func main() {

   if len(os.Args) < 2 {
      fmt.Printf("usage <filename> seed\n")
      return
   }

   seed, _  := strconv.Atoi(os.Args[1])
   rand.Seed(int64(seed))

   var  wg1,  wg2 sync.WaitGroup
   START()
   
   wg1.Add(2)
   go alpha(&wg1)
  
   wg2.Add(3)
   go beta(&wg2)
   go gamma(&wg2)
   go delta(&wg2)
   wg2.Wait()
   
   go epsilon(&wg1)
   wg1.Wait()
   
   STOP()
}
