package main


import (
    "testing"
    "reflect"
)


func TestCleanse(t *testing.T) {
    t.Log("Trying to preprocess text")
    raw := "И, однако, не чай! Или скажете? Я и не скажу."
    expectedText := "и однако не чай или скажете я и не скажу"
    removeChars := []string{",", ".", "?", "!"}
    if res := cleanse(raw, removeChars); res != expectedText {
        t.Errorf("Expected %s, got %s", expectedText, res)
    }
}


func TestTokenize(t *testing.T) {
    t.Log("Trying to split text to tokens")
    raw := "и однако не чай или скажете я и не скажу"
    expectedRes := []string{"и", "однако", "не", "чай", "или", "скажете", "я", "и", "не", "скажу"}
    res := tokenize(raw)
    if eq := reflect.DeepEqual(res, expectedRes); !eq {
        t.Errorf("Expected %v, got %v", expectedRes, res)
    }
}


func TestCount(t *testing.T) {
    t.Log("Trying to count tokens in slice")
    raw := []string{"бу", "му", "ку", "му", "бу", "бу"}
    expectedRes := map[string]int{"бу": 3, "му": 2, "ку": 1}
    res := countTokens(raw)
    if eq := reflect.DeepEqual(res, expectedRes); !eq {
        t.Errorf("Expected %v, got %v", expectedRes, res)
    }
}


func TestTopTokens(t *testing.T) {
    t.Log("Trying to find top tokens")
    raw := map[string]int{"бу": 3, "му": 2, "ку": 1, "вуу": 10}
    expectedRes := []token{
        {"вуу", 10},
        {"бу", 3},
    }
    res := topTokens(raw, 2)
    if eq := reflect.DeepEqual(res, expectedRes); !eq {
        t.Errorf("Expected %v, got %v", expectedRes, res)
    }
}


func TestWordCount(t *testing.T) {
    t.Log("Complete operation")
    raw := "У попа (ПОПА) была собака (СоБаКа), у попа, у попа!"
    expectedRes := []token{
        {"попа", 4},
        {"у", 3},
        {"собака", 2},
    }
    res := WordCount(raw, []string{",", "!", "(", ")"}, 3, false)
    if eq := reflect.DeepEqual(res, expectedRes); !eq {
        t.Errorf("Expected %v, got %v", expectedRes, res)
    }
}
