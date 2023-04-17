package sitemeta

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rajatxs/go-cosmic-echoes/types"
	"github.com/rajatxs/go-cosmic-echoes/util"
)

func GetSource(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, errors.New("failed to get source")
	}
	return response.Body, nil
}

func GetDocument(source io.ReadCloser) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(source)
}

// Extract title of the page from given document
func GetTitle(doc *goquery.Document) string {
	var (
		title string
		node  *goquery.Selection
	)

	if node = doc.Find("title"); node.Length() > 0 {
		// <title></title>
		title = node.First().Text()
	} else if node = doc.Find("meta[name='apple-mobile-web-app-title']"); node.Length() > 0 {
		// <meta name="apple-mobile-web-app-title" />
		title = node.First().AttrOr("content", "")
	} else if node = doc.Find("meta[property='og:title']"); node.Length() > 0 {
		// <meta property="og:title" />
		title = node.First().AttrOr("content", "")
	} else if node = doc.Find("meta[property='og:site_name']"); node.Length() > 0 {
		// <meta property="og:site_name" />
		title = node.First().AttrOr("content", "")
	} else if node = doc.Find("meta[name='application-name']"); node.Length() > 0 {
		// <meta name="application-name" />
		title = node.First().AttrOr("content", "")
	} else {
		title = ""
	}

	return strings.TrimSpace(title)
}

// Extract description of the page from given document
func GetDescription(doc *goquery.Document) string {
	var (
		desc string
		node *goquery.Selection
	)

	if node = doc.Find("meta[name='description']"); node.Length() > 0 {
		// <meta name="description" />
		desc = node.First().AttrOr("content", "")
	} else if node = doc.Find("meta[property='og:description']"); node.Length() > 0 {
		// <meta property="og:description" />
		desc = node.First().AttrOr("content", "")
	} else {
		desc = ""
	}

	return strings.TrimSpace(desc)
}

// Extract icon of the page from given document
func GetIcon(doc *goquery.Document, baseUrlStr string) string {
	var (
		icon string
		node *goquery.Selection
	)

	if node = doc.Find("link[rel='apple-touch-icon']"); node.Length() > 0 {
		// <link rel="apple-touch-icon" />
		icon = node.First().AttrOr("href", "")
	} else if node = doc.Find("link[rel='icon']"); node.Length() > 0 {
		// <link rel="icon" />
		icon = node.First().AttrOr("href", "")
	} else if node = doc.Find("link[rel='shortcut icon']"); node.Length() > 0 {
		// <link rel="shortcut icon" />
		icon = node.First().AttrOr("href", "")
	} else if node = doc.Find("link[rel='icon shortcut']"); node.Length() > 0 {
		// <link rel="icon shortcut" />
		icon = node.First().AttrOr("href", "")
	} else if node = doc.Find("link[type='image/x-icon']"); node.Length() > 0 {
		// <link type="image/x-icon" />
		icon = node.First().AttrOr("href", "")
	} else {
		icon = ""
	}

	if len(icon) > 0 {
		icon = util.GetAbsoluteUrl(baseUrlStr, icon)
	}

	return icon
}

func GetThumb(doc *goquery.Document) string {
	var (
		thumb string
		node  *goquery.Selection
	)

	if node = doc.Find("meta[property='og:image']"); node.Length() > 0 {
		// <meta property="og:image" />
		thumb = node.First().AttrOr("content", "")
	}

	return thumb
}

func GetSiteMetadata(url string, doc *goquery.Document) (*types.ResultSiteMeta, error) {
	var (
		meta = &types.ResultSiteMeta{}
		err  error
	)

	meta.Title = GetTitle(doc)
	meta.Description = GetDescription(doc)
	meta.Icon = GetIcon(doc, url)
	meta.Thumb = GetThumb(doc)
	return meta, err
}
