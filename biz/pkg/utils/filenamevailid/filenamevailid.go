package filenamevailid

import (
	"errors"
	"strings"
)

func Filenamevailid(filename, allowfiletype string) error {
	// log.Fatal(filename, allowfiletype)
	ext := filename[strings.LastIndex(filename, "."):]
	allowfilelist := strings.Split(allowfiletype, ",")
	//check ext without using utils
	isinallowallowfilelist := false
	for _, v := range allowfilelist {
		if v == ext {
			isinallowallowfilelist = true
			break
		}
	}
	if !isinallowallowfilelist {
		return errors.New("error file type")
	}
	return nil
}
