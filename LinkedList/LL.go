/* LL.go linked list */
package main

import ( "fmt"; "sync"; "strconv"; "os" )

type Node struct{
   data int64
   next *Node
}

type LinkedList struct{
   head *Node
   tail *Node
   size int64
}

func addHead(LL *LinkedList, data int64){
   node := new(Node)
   node.data = data
   node.next = LL.head
   if LL.head == nil {
      LL.tail = node
   }
   LL.head = node
   LL.size++
}

func addTail(LL *LinkedList, data int64, wg *sync.WaitGroup){
   node := new(Node)
   node.data = data
   node.next = nil
   if LL.head == nil {
      LL.head = node
      LL.tail = node
      LL.size++
      wg.Done()
      return
   }
   LL.tail.next = node
   LL.tail = node
   LL.size++
   wg.Done()
}

func printList(LL *LinkedList){
   var ptr *Node = LL.head
   for !(ptr==nil) {
      fmt.Printf("%3d", ptr.data)
      ptr = ptr.next
   }
   fmt.Println()
}

func makeHist(LL *LinkedList, n int64){
   sz := LL.size
   if sz != n {
      fmt.Printf("size (%2d) is not n (%2d)\n", sz, n)
   }
   theHist := make([]int64, n)
   var ptr *Node=LL.head
   for !(ptr==nil){
      theHist[ptr.data]++
      ptr = ptr.next
   }
   for i:=int64(0); i<sz; i++{
      if theHist[i] != 1 {
         fmt.Printf("%7d: %3d\n",i,theHist[i])
      }
   }
   for i:=int64(0); i<n; i++{
      fmt.Printf("%3d",i)
   }
   fmt.Println()
}

func main() {
   n, _ := strconv.ParseInt(os.Args[1], 10, 64)
   LL := new(LinkedList)
   var wg sync.WaitGroup
   
   wg.Add(int(n))
   for i := int64(0); i<n; i++ {
      go addTail(LL, i, &wg)
   }
   wg.Wait()

   makeHist(LL, n)
   printList(LL)
}
