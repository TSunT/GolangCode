package testreflect

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type JsonTimeFormat time.Time

func (formatTime *JsonTimeFormat) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*formatTime = JsonTimeFormat(now)
	return
}

func (formatTime JsonTimeFormat) MarshalJSON() (byt []byte, err error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(formatTime).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

type Movie struct {
	Title     string         `json:"title"`
	Price     float32        `json:"price"`
	Actors    []string       `json:"actors"`
	IssueDate JsonTimeFormat `json:"issueDate"`
}

func TestStructToJson() {
	movie := Movie{
		Title:     "小林子",
		Price:     10.98,
		Actors:    []string{"a", "b", "c", "d"},
		IssueDate: JsonTimeFormat(time.Now()),
	}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("convert to JSON Failed!")
	}

	fmt.Println("convert result: ", string(jsonStr))
}

func TestJsonToStruct() {
	// {"title":"123","price":10.98,"actors":["a","b","c","d"],"issueDate":"2024-08-08 09:22:12"}
	var movie Movie = Movie{}
	jsonStr := "{\"title\":\"小林子\",\"price\":10.98,\"actors\":[\"a\",\"b\",\"c\",\"d\"],\"issueDate\":\"2024-08-08 09:22:12\"}"
	err := json.Unmarshal([]byte(jsonStr), &movie)
	if err != nil {
		fmt.Println("unMarshal failed!")
	}
	fmt.Println("unMarshal result: ", movie)
}
