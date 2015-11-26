package main

import (
   "fmt"
   "unicode/utf8"
    "log"
   "os"
   "bufio"
   "io"
    "time"
)

type Text []byte
const (
    TOTAL_MATCH   = 0
    PREFIX_MATCH  = 1
    NOT_MATCH     = 2
)


type Node struct {
    Word    string
    Isword  bool
    Child   map[string]*Node
	size    int
}

func CreateTrie(data []string) Node{
     root := Node{Child : make(map[string]*Node)}
     for  _, txt := range data {
	    words := TexttoWord([]byte(txt))
	    root.createNode(words,0)
	 }
	 return root
}

func (node *Node) createNode(words []string,i int) {
   if i>=len(words) {
      return
   }

   
   n := lookup(words[i],node)
   if n != nil {
       if i < len(words)-1 {
	      i++
	      n.createNode(words,i)
	   }else if i == len(words)-1 {
	      n.Isword = true
	   }
   }
}

func lookup(word string,node *Node) *Node{
   n, ok := node.Child[word]
   if !ok {
     n = new(Node)
	 n.Word = word
	 n.Child = make(map[string]*Node)
	 node.Child[word] = n
	 node.size++
   }
   return n
}

func (node *Node) match(words []string,cursor int) int {
    if words == nil{
	   return NOT_MATCH
	}
	
	w, ok := node.Child[words[cursor]]
	if ok {
	   if cursor == len(words)-1 {
	      if w.Isword == true {
		      return TOTAL_MATCH
		  }
		  if w.size>0{
		      return PREFIX_MATCH
		  }
	   } else {
	      cursor++
	      return w.match(words,cursor)
	   }
	}
	return NOT_MATCH
}



func TexttoWord(text Text) []string {
    var output []string
    output = make([]string, 0)
    cursor := 0
    i := 0
    for cursor < len(text) {
      p, size := utf8.DecodeRune(text[cursor:])
      //fmt.Printf("%c,%v\n",p,size)
      output =append(output, string(p))
      i++
      cursor += size
    }
    return output
}





func init() {
   fmt.Printf("init\n")
}

func main() {
   //words := []string{"中国","美国","德国","A股H","意大利"}
   words := make([]string, 0)
   file, err := os.Open("main2012.dic")
   if err !=nil {
      fmt.Println("error")
	  return
   }
   defer file.Close()
   br := bufio.NewReader(file)

   for {
     line, _, err := br.ReadLine()
	 if err == io.EOF {
	    break
	 }
	 words = append(words,string(line))
   }
   t1 := time.Now()
   root := CreateTrie(words)
    t2 := time.Now()
    log.Println(t2.Sub(t1))
   log.Println("dictionary is ok")
   
   flag := root.match(TexttoWord([]byte("中国")),0)
   if flag==TOTAL_MATCH{
      fmt.Printf("total match\n")
   } else if flag==NOT_MATCH{
      fmt.Printf("not match\n")
   }else{
      fmt.Printf("prefix match\n")
   }
}
