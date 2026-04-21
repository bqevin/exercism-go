package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
    First WeekSchedule = iota
    Second
    Third
    Fourth
    Last
    Teenth
)

func (ws WeekSchedule) String() string {
    options := []string{"first", "second", "third", "fourth", "last", "teenth"};
    return options[ws];
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
    // create date : first day of month
    t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC);
    var min, max int
    switch wSched {
        // 1 - 7 
        case First:
        	min, max = 1, 7;
        case Second:
        	min, max = 8, 14;
    	case Third:
        	min, max = 15, 21;
        case Fourth:
            min, max = 22, 28;
        case Last:
        	lastDay := time.Date(year, month + 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1).Day();
        	min, max = lastDay - 6, lastDay;
        case Teenth:
            min, max = 13, 19;
        default:
        	min, max = 1, 31;
    }
    
    for i := min; i <= max; i++ {  
        cd := t.AddDate(0, 0, i - 1); // add 0 on first iteration
        // return exact day of month
        if cd.Weekday() == wDay { 
            return cd.Day(); 
        }
    }   
    return t.Day();
}
