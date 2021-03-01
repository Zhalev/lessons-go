package hw_4

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	getFile("../hw-1")
}

func getFile(pathy string) {
	dir, _ := path.Split(pathy)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == true {
			continue
		}
		fn := file.Name()
		b, _ := ioutil.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		val := string(b)
		if len(val) > 0 {
			os.Setenv(fn, val)
		}
	}
}
