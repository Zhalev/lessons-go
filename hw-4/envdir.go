package hw_4

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

var (
	e   string
	p   string
	arg []string
)

func init() {
	flag.StringVar(&e, "e", "", "directory environment files")
	flag.StringVar(&p, "p", "", "program or cmd")
	arg = flag.Args()
}

func main() {

	flag.Parse()

}

func readDir(dir string) (map[string]string, error) {
	ms := make(map[string]string)
	dir, _ = path.Split(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
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
			ms[fn] = val
			os.Setenv(fn, val)
		}
	}
	return ms, nil
}
func RunCmd(cmd []string, env map[string]string) int {

	exec.Command("prog")
	return 1

}
