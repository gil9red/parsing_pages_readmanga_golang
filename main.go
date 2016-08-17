package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func download(url string) {
	fmt.Println("Downloading " + url + " ...")
	rs, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer rs.Body.Close()

	d, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := string(d[:])

	// Parsing pages chapter
	re := regexp.MustCompile(`rm_h\.init\(.*?\[(.+)\].*?\);`)
	matchs := re.FindStringSubmatch(s)

	// Is match
	if len(matchs) > 0 {
		s := matchs[1]

		// Get list pages
		re = regexp.MustCompile(`\[.+?\]`)
		matchs = re.FindAllString(s, -1)
		fmt.Println("Number pages", len(matchs))

		fmt.Println("Pages:")
		for i, str := range matchs {
			// Page is string: ['auto/00/op','http://e1.postfact.ru/',"/v1ch3/OnePiece_Log01_Chapter003_01.png",1055,1520]

			// Removing unnecessary characters
			str = strings.Replace(str, "[", "", -1)
			str = strings.Replace(str, "]", "", -1)
			str = strings.Replace(str, "\"", "", -1)
			str = strings.Replace(str, "'", "", -1)

			// Get part url page
			parts := strings.Split(str, ",")
			url := parts[1] + parts[0] + parts[2]
			fmt.Printf("  %d: %s\n", i+1, url)
		}
	}
}

func main() {
	url := "http://readmanga.me/one__piece/vol1/3?mature="
	download(url)
}
