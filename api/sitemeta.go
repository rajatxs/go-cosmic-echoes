package handler

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rajatxs/go-cosmic-echoes/sitemeta"
	"github.com/rajatxs/go-cosmic-echoes/types"
	"github.com/rajatxs/go-cosmic-echoes/util"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		url    string
		err    error
		source io.ReadCloser
		doc    *goquery.Document
		meta   *types.ResultSiteMeta
	)

	if r.URL.Query().Has("url") {
		url = r.URL.Query().Get("url")
	} else {
		util.SendResponse(&w, 400, "url is missing in query parameters", nil)
		return
	}

	url = strings.ToLower(url)

	if source, err = sitemeta.GetSource(url); err != nil {
		log.Println(err)
		util.SendResponse(&w, 500, "failed to get source", nil)
		return
	} else {
		defer source.Close()
	}

	if doc, err = sitemeta.GetDocument(source); err != nil {
		log.Println(err)
		util.SendResponse(&w, 500, "failed to parse source", nil)
		return
	}

	if meta, err = sitemeta.GetSiteMetadata(url, doc); err != nil {
		log.Println(err)
		util.SendResponse(&w, 500, "failed to extract metadata", nil)
		return
	}

	if err = util.SendResponse(&w, 200, "Ok", meta); err != nil {
		log.Println(err)

		w.WriteHeader(500)
		w.Write([]byte("something went wrong"))
	}
}
