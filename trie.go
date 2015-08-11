package main

import (
   "fmt"
   "unicode/utf8"
)

type Text []byte

type Node struct {
    Word    string
    Isword  bool
    Child   []*Node
}

/*func CreateTrie(data *string) {

}

func (node *Node) GetOne(key string){
}

func (node *Node) DeleteOne(key string){
}

func (node *Node) InsertOne(key string){
}*/


func TexttoWord(text Text) []string {
    var output []string
    output = make([]string, len(text)/3)
    cursor := 0
    i := 0
    for cursor < len(text) {
      p, size := utf8.DecodeRune(text[cursor:])
      fmt.Printf("%c,%v\n",p,size)
      output[i] = string(p)
      i++
      cursor += size
    }
    return output
}


func main() {
   words := []string{"中国","美国","德国","法国","意大利"}
   for _, word := range words {
      fmt.Printf("%v\n",word)
      key := TexttoWord([]byte(word))
      fmt.Printf("%v\n",key)
   }
}
