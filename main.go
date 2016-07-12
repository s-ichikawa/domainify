package main

import (
    "math/rand"
    "time"
    "bufio"
    "os"
    "strings"
    "unicode"
    "fmt"
)

const allowedChars = "abcdefghijkl"

func main()  {
    rand.Seed(time.Now().UTC().UnixNano())
    s := bufio.NewScanner(os.Stdin)
    println("input top level domains.")
    var tlds = []string{}
    for s.Scan() {
        domains := strings.ToLower(s.Text())
        if (len(domains) == 0) {
            break
        }
        tlds = append(tlds, domains)
    }
    fmt.Println(tlds)
    for s.Scan() {
        text := strings.ToLower(s.Text())
        var newText []rune
        for _, r := range text {
            if unicode.IsSpace(r) {
                r = '-'
            }
            if !strings.ContainsRune(allowedChars, r) {
                continue
            }
            newText = append(newText, r)
        }
        fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
    }

}