/* bank.go */
package main

import ( "fmt"; "math/rand"; "os"; "strconv" )

type logEntry struct {
    accNum int
    amount int
}

func transactions(accounts[]int, theLog[] logEntry, logIdx *int, numAccs int, maxTrans int) {
   numTrans := rand.Intn(maxTrans)
   for i:=0; i<numTrans; i++{
      accNum := rand.Intn(int(numAccs))
      amount := rand.Intn(200)-100
      accounts[accNum] += amount
      theLog[*logIdx].accNum  = accNum
      theLog[*logIdx].amount  = amount
      *logIdx++
   }
}

func audit(theLog[ ] logEntry, logIdx int, expectedBal[ ] int){
   for i:=0; i<logIdx; i++{
      expectedBal[theLog[i].accNum] += theLog[i].amount
   }
}

func main() {
   numAccs, _  := strconv.Atoi(os.Args[1])
   seed, _     := strconv.ParseInt(os.Args[2], 10, 64)
   maxTrans, _ := strconv.Atoi(os.Args[3])
   accounts := make([]int, numAccs)
   theLog   := make([]logEntry, maxTrans*numAccs)
   logIdx   := 0
   rand.Seed(seed)

   transactions(accounts, theLog, &logIdx, numAccs, maxTrans)
   expectedBal := make([]int, numAccs)
   audit(theLog, logIdx, expectedBal)

   for i:=0; i<numAccs; i++ {
      
      fmt.Printf("acc no %2d balance %5d expected %5d", i, accounts[i], expectedBal[i]);
      
      if accounts[i] != expectedBal[i] {
         fmt.Printf(" failed")
      }
      fmt.Printf("\n")
   }
   fmt.Println("numAccs", numAccs, "logEntries", logIdx)
}
