package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`);
	return re.MatchString(text);
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[=*\-~$]*>`);
    return re.Split(text, -1);
}

func CountQuotedPasswords(lines []string) int {
    count := 0;
	re := regexp.MustCompile(`"(?i).*password.*"`);
    for _, v := range lines {
        count += len(re.FindAllString(v, -1));
    }
    return count;
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line(\d+)`);
    return re.ReplaceAllString(text, "");
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+([a-zA-Z]+\d+)`);
    for i, line := range lines {
        if re.FindString(line) != "" {
            lines[i] = fmt.Sprintf("[USR] %v %s", re.FindStringSubmatch(line)[1], line);
        } else {
            lines[i] = line;
        }
    }
    return lines;
}
