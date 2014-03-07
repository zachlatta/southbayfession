package misc

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func Prerender(res http.ResponseWriter, req *http.Request) {
	prerender := false

	uas := regexp.MustCompile(`(?i)(baiduspider|facebookexternalhit|twitterbot|rogerbot|linkedinbot|embedly|bufferbot|quora link preview)`)
	fileExtensions := regexp.MustCompile(`(?i)\.(js|css|xml|less|png|jpg|jpeg|gif|pdf|doc|txt|ico|rss|zip|mp3|rar|exe|wmv|doc|avi|ppt|mpg|mpeg|tif|wav|mov|psd|ai|xls|mp4|m4a|swf|dat|dmg|iso|flv|m4v|torrent)$`)

	if uas.MatchString(req.UserAgent()) {
		prerender = true
	}

	if regexp.MustCompile(`_escaped_fragment_`).MatchString(req.URL.String()) {
		prerender = true
	}

	if regexp.MustCompile(`Prerender`).MatchString(req.UserAgent()) {
		prerender = false
	}

	if fileExtensions.MatchString(req.URL.Path) {
		prerender = false
	}

	if prerender {
		basePrerenderUrl := os.Getenv("PRERENDER_URL")
		prerenderUrl := basePrerenderUrl + req.URL.String()

		resp, err := http.Get(prerenderUrl)
		defer resp.Body.Close()
		if err != nil {
			log.Println(err)
			return
		}

		io.Copy(res, resp.Body)
	}
}
