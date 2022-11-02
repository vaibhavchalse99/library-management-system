package books

import "time"

func getEndOfTheDay(t string) (err error, timeStamp time.Time) {
	timeStamp, err = time.Parse("2006-01-02", t)
	if err != nil {
		return
	}
	year, month, day := timeStamp.Date()
	timeStamp = time.Date(year, month, day, 23, 59, 59, 999999999, timeStamp.Location())
	return
}
