package wordcount

import (
    "regexp"
    "strings"
    "fmt"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := make(Frequency)
    // ignore: !&@$%^&,: except: '
    re := regexp.MustCompile(`[a-zA-Z0-9']+`);
    words := re.FindAllString(phrase, -1);
    fmt.Printf("words: %v \n", words);
    // text as strings.ToLow(key), count as value 
    for i, v := range words {
        // trim sorround space and apostrophe (accept only when between words like it's)
        v = strings.Trim(strings.ToLower(v), " '")
        // skip spaces
        if v != ""{
            f[v] = 1;
            for j, k := range words {
                // trim sorround space and apostrophe (accept only when between words like it's)
                k = strings.Trim(strings.ToLower(k), " '")
                // skip repeating count
                if i == j { continue; }
                if k == v { f[v] += 1; }
            }  
        }  
    }
    return f;
}
