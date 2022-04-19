package inputOutputGo

import (
	"bufio"
	"os"
)

func OpenFile() (id, pw, signpw string) {
	fo, err := os.Open("idpw.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	reader := bufio.NewReader(fo)
	for i:= 1; i<= 3; i++ {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		}
		switch (i) {
		case 1:
			id = string(line)
			break
		case 2:
			pw = string(line)
			break
		case 3:
			signpw = string(line)
			break
		}
	}
	return
}
