package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Esta función recibe una URL, se utiliza para obtener contenido HTML y extraerlo en texto plano.
func fetchAndExtractText(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var textBuilder strings.Builder
	doc.Find("p, h1, h2, h3, h4, h5, h6, li").Each(func(i int, s *goquery.Selection) {
		textBuilder.WriteString(s.Text())
		textBuilder.WriteString(" ")
	})

	// Quitamos espacios en blanco y saltos de línea
	text := strings.TrimSpace(textBuilder.String())
	text = strings.Join(strings.Fields(text), " ")

	return text, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la URL: ")
	url, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error leyendo la URL")
	}
	url = strings.TrimSpace(url)

	fmt.Println("Obteniendo HTML de: ", url)
	text, err := fetchAndExtractText(url)
	if err != nil {
		log.Fatalf("Error obteniendo texto y extrayendolo")
	}
	fmt.Println("Texto obtenido:")
	fmt.Println(text)
}
