package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func count(file *os.File, m map[string]map[string]int) {
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    key := scanner.Text()
    if _, ok := m[key]; !ok { // avoid having a nil map
      m[key] = make(map[string]int)
    }
    m[key][file.Name()]++
  }
}

func main() {
  var m = make(map[string]map[string]int)
  for _, filename := range os.Args[1:] {
    if file, err := os.Open(filename); err != nil {
      fmt.Println("Failed to open ", filename)
      continue
    } else {
      count(file, m)
      file.Close()
    }
  }

  for line, stats := range m {
    var sum int
    var fileList []string
    for fileName, cnt := range stats {
      sum += cnt
      fileList = append(fileList, fileName)
    }

    if sum > 1 {
      s := strings.Join(fileList, " ")
      fmt.Printf("%d %s %s\n", sum, s, line)
    }
  }
}
