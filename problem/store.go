package problem

import (
	"errors"
	"os"
	"path/filepath"
)

const dir = "./problems"

func join(name string) string {
	return filepath.Join(dir, name)
}

func readDir() ([]string, error) {
	ds, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	rs := make([]string, 0, len(ds))
	for _, v := range ds {
		rs = append(rs, v.Name())
	}
	return rs, nil
}

func readFile(name string) *Problem {
	buff, err := os.ReadFile(join(name))
	if err != nil {
		return nil
	}
	return Gob2Problem(buff)
}

func writeFile(name string, p Problem) error {
	buff := Problem2Gob(p)
	if buff == nil {
		return errors.New("cannot convert problem to gob")
	}
	if _, err := os.Stat(join(name)); os.IsNotExist(err) {
		if _, err := os.Create(join(name)); err != nil {
			return err
		}
	}
	if err := os.WriteFile(join(name), buff, 0660); err != nil {
		return err
	}
	return nil
}
