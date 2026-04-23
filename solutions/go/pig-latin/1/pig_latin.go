package piglatin

import (
    "strings"
    "regexp"
)

func Sentence(sentence string) string {
    var sb strings.Builder;
    re := regexp.MustCompile(`[a-zA-Z]+`);
    words := re.FindAllString(sentence, -1);

    for i, w := range words {
        l := strings.Trim(strings.ToLower(w)," ");
        
        // space out words if it's a sentence
        if i > 0 {
            sb.WriteByte(' ');
        }
        
        switch {
            // rule 1: if word idx 0 is vowel (a,e,i,o,u) or idx 0 & 1 are "xr" or "yt" 
            case isVowel(l[0]):
            	sb.WriteString(l);
            case (l[0] == 'x'  && l[1] == 'r') || (l[0] == 'y'  && l[1] == 't'): // xray or yttria
        		sb.WriteString(l);
            case !isVowel(l[0]):
            	// rule 3: if word idx 0, 1 is "qu" or idx 1, 2 is "qu", move every letter to end till 'u'
            	if l[0] == 'q' && l[1] == 'u' { // quick
                    sb.WriteString(l[2:]); 
                    sb.WriteString(l[:2]);
                } else if l[1] == 'q' && l[2] == 'u' { // square
                    sb.WriteString(l[3:]);
                    sb.WriteString(l[:3]);
                // rule 4: if word idx 1, 2 is 'y', move to end
                } else if l[1] == 'y' { // my
                    sb.WriteString(l[1:]);
                    sb.WriteByte(l[0]);
                } else if l[2] == 'y' { // rhythm
                    sb.WriteString(l[2:]);
                    sb.WriteString(l[:2]);
                // rule 2: if word idx 0 or idx 0, 1 is consonant, move to end
                } else if !isVowel(l[0]) && !isVowel(l[1]) && !isVowel(l[2]) { // school
                    sb.WriteString(l[3:]);
                    sb.WriteString(l[:3]);
                } else if !isVowel(l[0]) && !isVowel(l[1]) { // chair
                    sb.WriteString(l[2:]);
                    sb.WriteString(l[:2]);
                } else {
                   	sb.WriteString(l[1:]);
            		sb.WriteByte(l[0]); 
                } 	
        }
        // add "ay" at end
        sb.WriteString("ay");
    }
    
    
	return sb.String();
}

func isVowel(c byte) bool{
    return  c == 'a' ||  c == 'e' || c == 'i' || c == 'o' || c == 'u';
}