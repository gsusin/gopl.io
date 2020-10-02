// Exercício 4.12

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Comic struct {
	Number     int `json:"num"`
	Img        string
	Alt        string
	Transcript string
}

const ComicsURL = "https://xkcd.com/"

func DownloadComics() (map[int]Comic, error) {
	ComicsList := make(map[int]Comic)

	for n := 2300; n <= 2311; n++ {
		var downloadURL string
		downloadURL = ComicsURL + strconv.Itoa(n) + "/info.0.json"
		resp, err := http.Get(downloadURL)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("download failed: %s URL: %s", resp.Status, downloadURL)
		}
		var result Comic
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()
		ComicsList[n] = result
	}
	return ComicsList, nil
}

func CreateIndex(comicsList map[int]Comic) map[string]map[int]bool {
	index := make(map[string]map[int]bool)
	for _, comic := range comicsList {
		//Extrai termos
		fmt.Printf("Number: %d\tAlt: %s\n", comic.Number, comic.Alt)
		fields := strings.Fields(comic.Alt)
		//Adiciona termos aos mapas
		for _, field := range fields {
			if _, ok := index[field]; !ok {
				index[field] = make(map[int]bool)
			}
			index[field][comic.Number] = true
		}
	}
	return index
}

func main() {
	var comicsList map[int]Comic
	var err error
	comicsList, err = DownloadComics()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Baixados: %d\n", len(comicsList))
	index := CreateIndex(comicsList)
	fmt.Printf("Indexados: %d\n", len(index))
	//for _, c := range comicsList {
	//	fmt.Printf("Comic Number: %d\tComic Alt: %s\n", c.Number, c.Alt)
	//}
	fmt.Printf("Número de itens com '%s': %d\n", os.Args[1], len(index[os.Args[1]]))
	for comicNum := range index[os.Args[1]] {
		fmt.Printf("%d\t%s\n", comicNum, comicsList[comicNum].Alt)
	}

}
