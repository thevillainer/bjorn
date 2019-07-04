package bjorn

import (
  "regexp"
  "strings"
)

// IsAlpha:- true if string contains only alphabets
func IsAlpha(s string) bool {
  r, _ := regexp.MatchString(`(?i)^[a-z]+$`, s)
  return r
}

// IsUpper:- true if string contains only upper case alphabets
func IsUpper(s string) bool {
  r, _ := regexp.MatchString(`^[A-Z]+$`, s)
  return r
}

// IsLower:- true if string contains only lower case alphabets
func IsLower(s string) bool {
  r, _ := regexp.MatchString(`^[a-z]+$`, s)
  return r 
}

// IsDigit:- true if string contains only digits
func IsDigit(s string) bool {
  r, _ := regexp.MatchString(`^[0-9]+$`, s)
  return r
}

// IsSpace:- true if string contains only spaces
func IsSpace(s string) bool {
  r, _ := regexp.MatchString(`^\s+$`, s)
  return r
}

// IsPunctuation:- true if string contains only punctuation
func IsPunctuation(s string) bool {
  r, _ := regexp.MatchString(`[!"#$%&'()*+,-./:;<=>?@[\]^_{|}~\\"]`, s)
  return r
}

// IsEmpty:- true if string == "" 
func IsEmpty(s string) bool {
  if s == "" {
    return true
  } else {
    return false
  }
}



// RemoveDuplicate:- remove duplicate from string
func RemoveDuplicate(s string) string {
  keys := make(map[string] bool)

  result := []string{}
  for _, k := range strings.Split(s, "") {
    if _, val := keys[k]; !val {
      keys[k] = true
      result = append(result, k)
    }
  }

  return strings.Join(result, "")
}

// SliceToMap:- create map using two string slices
func SliceToMap(sl1, sl2 []string) map[string]string {
  result := make(map[string] string)
  for i := 0; i < len(sl1); i++ {
    result[sl1[i]] = sl2[i]
  }

  return result
}

// ReverseMap:- reverse map
func ReverseMap(m map[string]string) map[string]string {
  result := make(map[string]string, len(m))

  for k, v := range m {
    result[v] = k
  }

  return result
}


func MaintainCase(text, matchText string) string {
  re := regexp.MustCompile(`[^\w\s]|[0-9]|\s`)
  textSlice := strings.Split(re.ReplaceAllLiteralString(text, ""), "")

  var result strings.Builder
  for _, v := range strings.Split(matchText, "") {
    if IsAlpha(v) {
      if IsUpper(v) {
        result.WriteString(strings.ToUpper(textSlice[0]))
        textSlice = textSlice[1:]
      } else {
        result.WriteString(strings.ToLower(textSlice[0]))
        textSlice = textSlice[1:]
      }
    } else {
      result.WriteString(v)
    }
  }


  return result.String()
}


// Mod:- negative mod
func Mod(n,m int) int {
  return ((n % m) + m) % m
}


// Range:- stupid python like range
func Range(n int) []int {
  result := []int{}

  for i := 0; i < n; i++ {
    result = append(result, i)
  }

  return result
}
