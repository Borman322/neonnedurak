package habr

import (
	"fmt"

	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func HabrGo() string  {
	URL := fmt.Sprintf("https://habr.com/ru/all/")
	res, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	linkAll := doc.Find(".tm-article-snippet__title")
	link, _ := linkAll.Find("a").Attr("href")
	linkText, _ := linkAll.Find("span").Html()
	// text := linkAll.Text()
	neonGondon := doc.Find(".article-formatted-body")
	muha, _ := neonGondon.Find("p").Html()
	// fmt.Println(text)
	text := fmt.Sprintf(`<a href\=\"https://habr.com%s\">%s%s</a>`, link, linkText, muha)
	fmt.Println(link)
	fmt.Println(muha)
	return (text)
}
