// implicitPipelineSieve.go
// A concurrent prime sieve
// source: golang.org/doc/play/sieve.go
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// modified to (1) take n from command line
//             (2) measure execution time

package main
import ( "fmt"; "os"; "strconv"; "time")

func Generate(ch chan<- int) {
   for i := 2; ; i++ {
      ch <- i
   }
}

func Filter(in <-chan int, out chan<- int, prime int) {
   for {
      i := <-in
      if i%prime != 0 {
         out <- i
      }
   }
}

func main() {
   n, _ := strconv.Atoi(os.Args[1])
   start := time.Now()
   ch := make(chan int)
   go Generate(ch)
   count := 0
   for i := 0; i < n; i++ {
      prime := <-ch
      //fmt.Println(prime)
      ch1 := make(chan int)
      go Filter(ch, ch1, prime)
      ch = ch1
      count++
   }
   elapse := time.Since(start)
   fmt.Println(count, elapse.Microseconds())
}
