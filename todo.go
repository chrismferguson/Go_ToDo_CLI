package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task string 
	Done bool 
	CreatedAt time.Time
	CompletedAt time.Time
}

type List []item 

// string prints out a formatted list 
// implements the fmt.Stringer interface 
// This is a naive implementation that prints out all items, prefixed by an order number and an X if the item is completed.
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}

		// Adjust the item number K to print numbers starting from 1 to 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}

func (l *List) Add(task string) {
	t := item {
		Task: task, 
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// complete method marks a ToDo item as completed bye 
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l 
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	} 
	//adjusting index for 0 based index 
	ls[i-1].Done = true 
	ls[i-1].CompletedAt = time.Now()

	return nil 
}

func (l *List) Delete(i int) error {
	ls := *l 
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	// adjusting index for 0 based index 
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err 
	}
	return os.WriteFile(filename, js, 0644)
}

// get method opens the file, decodes the json data and parses it into a list 
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist){
			return nil 
		}
		return err 
	}
	if len(file) == 0 {
		return nil 
	}
	return json.Unmarshal(file, l) 
}

