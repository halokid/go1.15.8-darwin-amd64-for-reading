package main

import (
  "bytes"
  "fmt"
  "log"
)

// todo: how to make a bits slice example!
// todo: the element of the slice express as `displament integer`, like `77` left
// todo: move 3 location, it's  `77 <<< 3`, taht will be a unit number store in
// todo: the bits slice, so you can not change the `0000 0010` to `0000 1010` directly
// todo: you need to use some agorithm to express the bit slice to uint like this example code

type IntSet struct {
  words  []uint64
}

func (s *IntSet) Add(x int) {
  word, bit := x/64, uint(x % 64)
  log.Printf("Add() bit -->>> %+v", bit)
  for word >= len(s.words) {
    s.words = append(s.words, 0)
  }
  log.Printf("Add() (1 << bit) -->>> %+v", 1 << bit)
  s.words[word] |= 1 << bit
}

func (s *IntSet) Has(x int) bool {
  word, bit := x/64, uint(x%64)
  log.Printf("word -->>> %+v, bit -->>> %+v", word, bit)

  a := s.words[word]
  log.Printf("s.words[word] -->>> %+v", a)
  b := 1 << bit
  log.Printf("1 << bit -->>> %+v", b)
  retCondition := (s.words[word] & (1 << bit)) != 0
  //retCondition :=  a & b

  log.Printf("retCondition -->>> %+v", retCondition)
  //return word < len(s.words) && (s.words[word] & (1 << bit) != 0)
  return word < len(s.words) && retCondition
}

func (s *IntSet) UnionWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] |= tword
    } else {
      s.words = append(s.words, tword)
    }
  }
}

func (s *IntSet) String() string {
  var buf bytes.Buffer
  buf.WriteByte('{')
  for i, word := range s.words {
    if word == 0 {
      continue
    }
    for j := 0; j < 64; j++ {
      if word & (1 << uint(j)) != 0 {
        if buf.Len() > len("{") {
          buf.WriteByte(' ')
        }
        fmt.Fprintf(&buf, "%d", 64 * i + j)
      }
    }
  }
  buf.WriteByte('}')
  return buf.String()
}


// test -------------------------------------------
func TestAdd() {
  is := IntSet{
    words:  make([]uint64, 999),
  }
  is.Add(177)
  log.Printf("is -->>> %+v", is)
  log.Println(is.String())
}

func TestHas() {
  is := IntSet{
    words:  make([]uint64, 999),
  }
  log.Printf("len, cap of is.words -->>> %+v, %+v, is.words[0] is -->>> %+v",
    len(is.words), cap(is.words), is.words[0])
  b := is.Has(99)
  log.Printf("b -->>> %+v", b)
}

func main() {
  TestAdd()
}






