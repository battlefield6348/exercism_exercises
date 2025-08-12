package isogram

import (
    "strings"
)

func IsIsogram(word string) bool {
    existList := make(map[byte]bool)
    word = strings.ToLower(word)
    for i := range word {
        if word[i] == '-' || word[i] == ' '{
            continue
        }
        if existList[word[i]] {
            return false
        } else {
            existList[word[i]] = true
        }
    }
    return true
}
