package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
	"strings"
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
	if len(e) > 0 {
		envy := make(map[string]string)
		envy, err := ReadDir(e)
		if err != nil {
			log.Fatal(err)
		}
		if len(p) > 0 {
			RunCmd(p, arg, envy)
		}
	}
}

func ReadDir(dir string) (map[string]string, error) {
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
		b, _ := ioutil.ReadFile(path.Join(dir, fn))
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		lines := strings.Split(string(b), "\n")
		ls := strings.LastIndex(fn, ".")
		fn = string([]rune(fn)[:ls])
		if len(lines[0]) > 0 {
			ms[fn] = lines[0]
			//os.Setenv(fn, val)
		}
	}
	return ms, nil
}
func RunCmd(prog string, args []string, env map[string]string) int {
	//str1 := fmt.Errorf("%s/n" ,"Ашипка" )
	//fmt.Println(str1)
	//return 1
	i := 0
	cmd := exec.Command(prog)
	cmd.Args = args
	envy := []string{}
	for k, en := range env {
		fmt.Println(k, en)
		envy = append(envy, k+"="+en)
		i++
	}
	cmd.Env = envy
	std, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", std)
	return i

}
