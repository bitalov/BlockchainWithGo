package main

import (
  "bufio"
  "crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"time"
  "strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct{
  Index int // Position of data record in the block chain
  Timestamp string // automatically generated when the data is written
  BPM int // example of data the can be recorded here it's heartbeats per minute !
  Hash string // SHA256 identifier representing current data record
  PrevHash string // SHA256 ide.. for previous record
}

var Blockchain []Block // Represeting a Blockchain as a Slice.

func CalcHashForBlock (block Block) string {
  record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash

Hash := sha256.New() 
Hash.Write([]byte(record))
Hashed := Hash.Sum(nil)

return hex.EncodeToString(Hashed)
  
}

func CreateBlock(oldBlock Block , BPM int) (Block , error) {


  var newBlock Block 
  t := time.Now() 

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String() 
  newBlock.BPM = BPM 
  newBlock.PrevHash = oldBlock.Hash 
  newBlock.Hash = CalcHashForBlock(newBlock)

  return newBlock , nil
  
}

func isBlockValid(newBlock, oldBlock Block) bool {
  ExpectedIndex := oldBlock.Index + 1
  if ExpectedIndex != newBlock.Index {
    return false
  } else if oldBlock.Hash != newBlock.PrevHash {
    return false
  } else if CalcHashForBlock(newBlock) != 
   newBlock.Hash {
    return false
  } else {
     return true
  }
}

/* Classic Blockchain Issue 
// In case two nodes of our blockchain ecosystem both added blocks to their chains and we received them both. */
/*Solution
 We Pick the longest chain because it is the most up to date and have the latest blocks !!*/
/* Implementation */
func replaceChain(newBlocks []Block) {
  if len(newBlocks) > len(Blockchain) {
    Blockchain = newBlocks
  }
}

var BCserver chan []Block 
//Blockchain server


func main() {
  err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}


  BCserver = make(chan []Block)

  t := time.Now()
  gensisBlock := Block{0, t.String(), 0, "" , ""}
  spew.Dump(gensisBlock)
  
}
