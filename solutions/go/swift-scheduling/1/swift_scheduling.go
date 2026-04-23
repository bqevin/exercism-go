package swiftscheduling

import (
    "time"
    "strconv"
    "strings"
)

func DeliveryDate(start, delivery string) string {
    var cd time.Time;
    t, _ := time.Parse("2006-01-02T15:04:05", start);
    // workday : t = Mon, Tue, Wed, Thur, Fri
    // quarter: t = Jan - Mar, Apr - Jun, Jul - Sep, Oct - Dec
    // now = 2h;
    // asap = t > 13h; tomorrow at 13h
    // asap = t <= 13; today at 17h
    // eow = t.Weekday == Mon, Tue, Wed; Fri at 17h
    // eow = t.Weekday == Thur, Fri; Sun at 20h
    // t < nth month of year; t.Month() + n at 08h in first (workday)
    // t >= nth month of year; t.Year() + 1 on n.Month() at 08h in first (workday:Mon)
    // t <= nth quarter of year; on last n.Month() at 08h in last (workday:Fri)
    // t > nth of year; t.Year() + 1 on last n.Month() at 08h in last (workday:Fri)

    switch {
        case delivery == "NOW":
        	cd = t.Add(2 * time.Hour);
        case delivery == "ASAP":
        	// before 13h
        	if t.Hour() < 13 {
                cd = t.Add(time.Duration(17 - t.Hour()) * time.Hour); // 17h same day
            // at OR after 13h
            } else {
                cd = t.AddDate(0, 0, 1).Add(time.Duration(13 - t.Hour()) * time.Hour); // 13h next day
            }
        case delivery == "EOW":
        	// before Thur
        	if t.Weekday() == 1 || t.Weekday() == 2 || t.Weekday() == 3 {
                cd = t.AddDate(0, 0, int(5 - t.Weekday())).Add(time.Duration(17 - t.Hour()) * time.Hour); // 17h on Fri
            // after Thur
            } else {
                 cd = t.AddDate(0, 0, int(7 - t.Weekday())).Add(time.Duration(20 - t.Hour()) * time.Hour); // 20h on Sun
            }
        case strings.HasSuffix(delivery, "M"):
        	numPart := delivery[:len(delivery) - 1];
        	m, _ := strconv.Atoi(numPart);
        	if t.Month() < time.Month(m) {
               cd = time.Date(t.Year(), time.Month(m), 1, 8, 0, 0, 0, time.UTC); 
            } else {
                cd = time.Date(t.Year(), time.Month(m), 1, 8, 0, 0, 0, time.UTC).AddDate(1, 0, 0); 
                if cd.Weekday() > 5 {
                    ncd := cd;
                    cd = ncd.AddDate(0, 0, int(7 - ncd.Weekday() + 1));
                }
            }
        	
        case len(delivery) == 2 && delivery[0] == 'Q':
        	q := delivery[1] - '0'; // = decimal ASCII value for 0
            // q1 = mar-4, q2 = jun-7, q3 = sep-10, q4 = dec-13
        	endQ := time.Month((3 * q) + 1);
        	if t.Month() < time.Month(endQ) {
                cd = time.Date(t.Year(), time.Month(endQ), 0, 8, 0, 0, 0, time.UTC); 
            } else {
                cd = time.Date(t.Year(), time.Month(endQ), 0, 8, 0, 0, 0, time.UTC).AddDate(1, 0, 0); 
            }
        
            if cd.Weekday() > 5 {
                ncd := cd;
                cd = ncd.AddDate(0, 0, -int(ncd.Weekday() - 5));
            }
    }
    
	return cd.Format("2006-01-02T15:00:00");
}
