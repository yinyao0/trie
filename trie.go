package main

import (
   "fmt"
   "unicode/utf8"
   "unicode"
   "os"
   "bufio"
   "io"
)

type Text []byte


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
	   return 2
	}
	
	w, ok := node.Child[words[cursor]]
	if ok {
	   if cursor == len(words)-1 {
	      if w.Isword == true {
		      return 0
		  }
		  if w.size>0{
		      return 1
		  }
	   } else {
	      cursor++
	      return w.match(words,cursor)
	   }
	}
	return 2
}


func TexttoWord(text Text) []string {
    var output []string
	tmp := make([]Text, len(text))
    output = make([]string, 0)
	start :=0
    cursor := 0
	currentWord := 0
	inAlphanumeric := true
    i := 0
    for cursor < len(text) {
      p, size := utf8.DecodeRune(text[cursor:])
	  if size <= 2 && (unicode.IsLetter(p) || unicode.IsNumber(p)) {
	  		if !inAlphanumeric {
				start = cursor
				inAlphanumeric = true
			}
	  }else{
	  		if inAlphanumeric {
				inAlphanumeric = false
				if cursor != 0 {
					tmp[currentWord] = text[start:cursor]
					currentWord++
				}
			}
			tmp[currentWord] = text[cursor : cursor+size]
			currentWord++
	 }
       cursor += size
    }
	for r, _ := range tmp {
	    output = append(output,string(r))
		i++
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
   
   root := CreateTrie(words)
   fmt.Println("dictionary is ok")
   
   flag := root.match(TexttoWord([]byte("中国")),0)
   if flag==0{
      fmt.Printf("total match\n")
   } else if flag==2{
      fmt.Printf("not match\n")
   }else{
      fmt.Printf("prefix match\n")
   }
}
