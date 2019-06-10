package main


import (
    "fmt"
    "strings"
    "sort"
    "os"
    "log"
    "io/ioutil"
)


type token struct {
    Token string
    Count int
}


func cleanse(s string, remove []string) string {
    cleanedString := s
    for _, char := range remove {
        cleanedString = strings.Replace(cleanedString, char, "", -1)
    }
    return strings.ToLower(cleanedString)
}


func tokenize(s string) []string {
    return strings.Fields(s)
}


func countTokens(tokens []string) map[string]int {
    counts := map[string]int{}
    for _, token := range tokens {
        counts[token]++
    }
    return counts
}


func topTokens(counts map[string]int, n int) []token {
    var tokens []token
    for k, v := range counts {
        tokens = append(tokens, token{k, v})
    }

    sort.Slice(tokens, func(i, j int) bool {
        return tokens[i].Count > tokens[j].Count
    })

    return tokens[0:n]
}


func WordCount(text string, remove []string, top int, print bool) []token {
    cleanedText := cleanse(text, remove)
    tokens := tokenize(cleanedText)
    counts := countTokens(tokens)
    topCounts := topTokens(counts, top)
    if print {
        for _, token := range topCounts {
            fmt.Println(token.Token, ":", token.Count)
        }
    }
    return topCounts
}


func main() {
    file, err := os.Open("text.txt") // Voyna i Mir, first book
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    text, _ := ioutil.ReadAll(file)

    removeChars := []string{",", ".", "!", "?", ";", "—", "[", "]", "(", ")"}
    WordCount(string(text), removeChars, 10, true)
}
