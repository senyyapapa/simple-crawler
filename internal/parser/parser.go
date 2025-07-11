package parser

import (
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(body io.Reader, base string) ([]string, error) {
	var links []string
	baseUrl, err := url.Parse(base)

	if err != nil {
		return nil, err
	}

	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return links, err
			}
			return nil, z.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						refUrl, err := url.Parse(a.Val)
						if err != nil {
							continue
						}
						absoluteUrl := baseUrl.ResolveReference(refUrl).String()
						if strings.HasPrefix(absoluteUrl, "http") {
							links = append(links, strings.Split(absoluteUrl, "#")[0])
						}
						break
					}
				}
			}
		}
	}
}
