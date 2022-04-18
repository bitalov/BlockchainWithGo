package main

import (
  "crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
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

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

type Message struct {
	BPM int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := CreateBlock(Blockchain[len(Blockchain)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("PORT")
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
  err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	log.Fatal(run())
}
