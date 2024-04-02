package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	var inputURL string

	typeList := Menu()

	fmt.Print("Please input your url : ")
	_, err := fmt.Scanln(&inputURL)

	if err != nil {
		fmt.Println(err)
	}

	c := colly.NewCollector()

	if typeList == 1 {
		c.OnHTML("img[src]", func(e *colly.HTMLElement) {
			resultString := strings.Split(e.Attr("src"), "/")

			e.Request.Visit(e.Attr("src"))

			fmt.Println(e.Attr("src"))
			fmt.Println(resultString)

			err := DownloadFile("/img/"+resultString[len(resultString)-1], inputURL+e.Attr("src"))

			if err != nil {
				DownloadFile("/img/"+resultString[len(resultString)-1], inputURL+e.Attr("src"))
			}
		})
	} else if typeList == 2 {
		c.OnHTML("script[src]", func(e *colly.HTMLElement) {
			resultString := strings.Split(e.Attr("src"), "/")

			e.Request.Visit(e.Attr("src"))

			err := DownloadFile("/js/"+resultString[len(resultString)-1], inputURL+e.Attr("src"))

			if err != nil {
				DownloadFile("/js/"+resultString[len(resultString)-1], inputURL+e.Attr("src"))
			}
		})
	} else if typeList == 3 {
		c.OnHTML("*", func(e *colly.HTMLElement) {
			e.Request.Visit(e.Attr("src"))
		})

		DownloadFile("index.html", inputURL)
	}

	// On every a script element which has src attribute call callback
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// c.Visit("http://go-colly.org/")
	c.Visit(inputURL)
}

func DownloadFile(filePath string, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	osFile, err := os.Create("./scrapped-file/" + filePath)

	if err != nil {
		fmt.Println(err)
	}

	defer osFile.Close()

	_, err = io.Copy(osFile, resp.Body)

	return err
}

func Menu() int {
	var inputType int

	fmt.Println("---------MENU--------")
	fmt.Println("1. Images")
	fmt.Println("2. Scripts")
	fmt.Println("3. Alls")
	fmt.Println("---------------------")

	fmt.Print("Please choose your type files want to download : ")
	_, err := fmt.Scanln(&inputType)

	if err != nil {
		fmt.Println(err)
	}

	return inputType
}
