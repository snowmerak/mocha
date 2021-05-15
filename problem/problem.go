package problem

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Case struct {
	Input  []string
	Output []string
}

type Problem struct {
	Name     string `json:"name"`
	Conetent string `json:"content"`
	Cases    []Case `json:"cases"`
}

func init() {
	rand.Seed(time.Now().UnixNano())

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0770); err != nil {
			log.Fatal(err)
		}
	}
}

func SelectOne() *Problem {
	list, err := readDir()
	if err != nil {
		return nil
	}
	if len(list) == 0 {
		return nil
	}
	i := rand.Int() % len(list)
	return readFile(list[i])
}

func Select(name string) *Problem {
	return readFile(name)
}

func Enroll(str string) error {
	p := String2Problem(str)
	if p == nil {
		return errors.New("cannot convert string to problem")
	}
	return writeFile(p.Name, *p)
}

func Submit(str string) (string, error) {
	name, lang, code := String2Solution(str)
	p := readFile(name)
	if p == nil {
		return "", errors.New("problem name is invalid")
	}

	dt, err := time.Duration(0), error(nil)
	switch strings.ToLower(lang) {
	case "go":
		dt, err = evalGo(p.Cases, code)
	case "js":
		dt, err = evalJS(p.Cases, code)
	default:
		return "", errors.New("invalid language")
	}

	return dt.String(), err
}
