package models

import (
	"strconv"
	"time"
)

// PlayList --
var (
	PlayList map[string]*Track
)

func init() {
	PlayList = make(map[string]*Track)
	// u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	// UserList["user_11111"] = &u
}

type PL struct {
	Tracks []string
}

// Track --
type Track struct {
	ID     string
	Title  string
	Author []string
}

// AddTrack --
func AddTrack(t []Track, tt Track) []Track {
	tt.ID = "track_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	t = append(t, tt)
	return t
}

// // GetTrack --
// func GetUser(uid string) (t *Track, err error) {
// 	if t, ok := PlayList[uid]; ok {
// 		return t, nil
// 	}
// 	return nil, errors.New("Track not exists")
// }

// GetAllTrack --
func GetAllTrack() map[string]*Track {
	return PlayList
}
