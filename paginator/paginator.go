package paginator

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Paginator struct {
	simple      bool
	firstRow    int64
	currentPage int64
	lastPage    int64
	listRows    int64
	totalPage   int64
	total       int64
	hasMore     bool
	path        string
	fragment    string
	query       string
	varPage     string
}

//
func NewPaginator(listRows int64, currentPage int64, total int64) Paginator {
	totalPage := total / listRows
	if total%listRows > 0 {
		totalPage = total/listRows + 1
	}

	newPaginator := Paginator{
		simple:      false,
		currentPage: currentPage,
		listRows:    listRows,
		total:       total,
		totalPage:   totalPage,
		lastPage:    totalPage,
	}

	newPaginator.path = "/"

	newPaginator.firstRow = (newPaginator.currentPage - 1) * newPaginator.listRows

	if newPaginator.firstRow < 0 {
		newPaginator.firstRow = 0
	}

	return newPaginator
}

func (this *Paginator) Simple() bool {

	return this.simple
}

func (this *Paginator) FirstRow() int64 {
	return this.firstRow
}

func (this *Paginator) CurrentPage() int64 {
	return this.currentPage
}

func (this *Paginator) LastPage() int64 {
	return this.lastPage
}

func (this *Paginator) ListRows() int64 {
	return this.listRows
}

func (this *Paginator) TotalPage() int64 {
	return this.totalPage
}

func (this *Paginator) Total() int64 {
	return this.total
}

func (this *Paginator) HasMore() bool {
	return this.hasMore
}

func (this *Paginator) Url(page int64) string {
	if page <= 0 {
		page = 1
	}

	pageStr := strconv.FormatInt(page, 10)

	params, _ := url.ParseQuery(this.query)

	path := this.path

	if strings.Index(this.path, "[PAGE]") < 0 {

		params.Add(this.varPage, pageStr)

	} else {

		path = strings.Replace(this.path, "[PAGE]", pageStr, -1)

	}

	url := path
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return url
}

func (this *Paginator) GetUrlRange(start, end int64) map[int64]string {
	urls := make(map[int64]string, end-start+1)

	for page := start; page <= end; page++ {
		urls[page] = this.Url(page)
	}

	return urls
}

func init() {

	params, _ := url.ParseQuery("admin=1&url=1")
	fmt.Println(params.Encode())
}

func (this *Paginator) getPreviousButton(text string) string {
	if text == "" {
		text = "&laquo"
	}

	if this.CurrentPage() <= 1 {
		return this.getDisabledTextWrapper(text)
	}

	prevPage := this.CurrentPage() - 1
	url := this.Url(prevPage)

	return this.getPageLinkWrapper(url, prevPage)
}

func (this *Paginator) getNextButton(text string) string {
	if text == "" {
		text = "&raquo"
	}

	if false {
		return this.getDisabledTextWrapper(text)
	}

	nextPage := this.CurrentPage() + 1

	url := this.Url(nextPage)

	return this.getPageLinkWrapper(url, nextPage)
}

func (this *Paginator) getLinks() string {

	if this.Simple() {
		return ""
	}

	html := ""
	block := map[string]map[int64]string{
		"first":  nil,
		"slider": nil,
		"last":   nil,
	}

	var side int64 = 3
	var window int64 = side * 2

	if this.LastPage() < window+6 {
		block["first"] = this.GetUrlRange(1, this.LastPage())
	} else if this.CurrentPage() <= window {
		block["first"] = this.GetUrlRange(1, window+2)
		block["last"] = this.GetUrlRange(this.LastPage()-1, this.LastPage())
	} else if this.CurrentPage() > (this.LastPage() - window) {
		block["first"] = this.GetUrlRange(1, 2)
		block["last"] = this.GetUrlRange(this.LastPage()-(window+2), this.LastPage())
	} else {
		block["first"] = this.GetUrlRange(1, 2)
		block["slider"] = this.GetUrlRange(this.CurrentPage()-side, this.CurrentPage()+side)
		block["last"] = this.GetUrlRange(this.LastPage()-1, this.LastPage())
	}

	if block["first"] != nil {
		html += this.getUrlLinks(block["first"])
	}

	if block["slider"] != nil {
		html += this.getDots()
		html += this.getUrlLinks(block["slider"])
	}

	if block["last"] != nil {
		html += this.getDots()
		html += this.getUrlLinks(block["last"])
	}

	return html

}

func (this *Paginator) getAvailablePageWrapper(url, page string) string {

	return "<li><a href=" + url + ">" + page + "</a></li>"
}

func (this *Paginator) getDisabledTextWrapper(text string) string {

	return "<li class=\"active\"><span>" + text + "</span></li>"
}

func (this *Paginator) getDots() string {

	return this.getDisabledTextWrapper("...")
}

func (this *Paginator) getActivePageWrapper(text string) string {

	return "<li class=\"active\"><span>" + text + "</span></li>"
}

func (this *Paginator) Render() string {

	if this.Simple() {
		return fmt.Sprintf(
			"<ul class=\"pager\">%s %s</ul>",
			this.getPreviousButton(""),
			this.getNextButton(""))
	} else {
		return fmt.Sprintf(
			"<ul class=\"pagination\">%s %s %s</ul>",
			this.getPreviousButton(""),
			this.getLinks(),
			this.getNextButton(""))
	}
}

func (this *Paginator) getUrlLinks(urls map[int64]string) string {
	html := ""

	for page, url := range urls {
		html += this.getPageLinkWrapper(url, page)
	}

	return html
}

func (this *Paginator) getPageLinkWrapper(url string, page int64) string {
	pageStr := strconv.FormatInt(page, 10)
	if page == this.CurrentPage() {
		return this.getActivePageWrapper(pageStr)
	}

	return this.getAvailablePageWrapper(url, pageStr)
}
