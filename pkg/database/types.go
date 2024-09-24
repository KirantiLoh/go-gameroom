package database

import (
	"strings"
	"time"
)

type JSONTime time.Time

func (c *JSONTime) UnmarshalJSON(b []byte) error {
    value := strings.Trim(string(b), `"`) //get rid of "
    if value == "" || value == "null" {
        return nil
    }

    t, err := time.Parse("2004-11-19", value) //parse time
    if err != nil {
        return err
    }
    *c = JSONTime(t) //set result using the pointer
    return nil
}

func (c JSONTime) MarshalJSON() ([]byte, error) {
    return []byte(`"` + time.Time(c).Format("2004-11-19") + `"`), nil
}
