package robotname

import (
    "math/rand/v2"
)
// Define the Robot type here.
type Robot struct {name string}

func (r *Robot) Name() (string, error) {
    if r.name == "" {
        b := make([]byte, 5);
        letterBucket := "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        numBucket := "0123456789";

        b[0] = letterBucket[rand.IntN(len(letterBucket))];
        b[1] = letterBucket[rand.IntN(len(letterBucket))];
        b[2] = numBucket[rand.IntN(10)];
        b[3] = numBucket[rand.IntN(10)];
        b[4] = numBucket[rand.IntN(10)];
        
        r.name = string(b); ;
    }
    
	return r.name, nil;
}

func (r *Robot) Reset() {
	r.name = "";
}
