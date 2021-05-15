package problem

import (
	"bytes"
	"encoding/gob"
	"strings"
)

func Gob2Problem(buff []byte) *Problem {
	p := new(Problem)
	decoder := gob.NewDecoder(bytes.NewBuffer(buff))
	if decoder.Decode(p) != nil {
		return nil
	}
	return p
}

func Problem2Gob(p Problem) []byte {
	buff := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buff)
	if encoder.Encode(p) != nil {
		return nil
	}
	return buff.Bytes()
}

func String2Problem(str string) *Problem {
	p := new(Problem)
	datas := strings.Split(str, "\n")
	sw := 0
	for _, v := range datas {
		switch v {
		case "이름":
			fallthrough
		case "name":
			sw = 1
		case "설명":
			fallthrough
		case "explain":
			fallthrough
		case "explanation":
			sw = 2
		case "테스트 케이스":
			fallthrough
		case "cases":
			fallthrough
		case "test cases":
			sw = 3
		default:
			switch sw {
			case 1:
				p.Name += v
			case 2:
				p.Conetent += v
			case 3:
				p.Cases = append(p.Cases, Case{Input: strings.Split(v, " "), Output: nil})
				sw = 4
			case 4:
				p.Cases[len(p.Cases)-1].Output = strings.Split(v, " ")
				sw = 3
			}
		}
	}
	return p
}

func String2Solution(str string) (string, string, string) {
	sp := strings.SplitN(str, "\n", 3)
	if len(sp) < 3 {
		return "", "", ""
	}
	return sp[0], sp[1], sp[2]
}
