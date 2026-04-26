package raindrops

import (
    "strings"
    "strconv" 
)

func Convert(number int) string {
    var sb strings.Builder
    if number % 3 == 0 {
        sb.WriteString("Pling")
    } 
    if number % 5 == 0 {
        sb.WriteString("Plang")
    } 
    if number % 7 == 0 {
        sb.WriteString("Plong")
    } 

    if number % 5 != 0 && number % 3 != 0 && number % 7 != 0 {
        sb.WriteString(strconv.Itoa(number))
    }

    return sb.String()
}
