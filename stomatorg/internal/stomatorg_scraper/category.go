package stomatorg_scraper

import (
	"github.com/gocolly/colly/v2"
	"net/http"
	"strings"
)

type category struct {
	name               string
	parentCategoryName string
	link               string
}

func parseCategories() ([]category, error) {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))

	c := colly.NewCollector()
	c.WithTransport(t)
	// Set custom headers for all requests
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		r.Headers.Set("Accept-Language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Pragma", "no-cache")
		r.Headers.Set("Referer", "https://stomatorg.ru/")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-Site", "same-origin")
		r.Headers.Set("Sec-Fetch-User", "?1")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 YaBrowser/25.12.0.0 Safari/537.36")
		r.Headers.Set("sec-ch-ua", "\"Chromium\";v=\"142\", \"YaBrowser\";v=\"25.12\", \"Not_A Brand\";v=\"99\", \"Yowser\";v=\"2.5\"")
		r.Headers.Set("sec-ch-ua-mobile", "?0")
		r.Headers.Set("sec-ch-ua-platform", "\"Linux\"")
		r.Headers.Set("Cookie", "PHPSESSID=oLniOxDVKWiQojhyXPpJTTPMLG0kiiRj; BITRIX_SM_GUEST_ID=34065829; mindboxDeviceUUID=c26a2121-46f6-412e-aa7e-cdcba7acba2e; directCrm-session=%7B%22deviceGuid%22%3A%22c26a2121-46f6-412e-aa7e-cdcba7acba2e%22%7D; _ym_uid=1770832266394185598; _ym_d=1770832266; roistat_visit=7247016; roistat_first_visit=7247016; roistat_visit_cookie_expire=1209600; roistat_is_need_listen_requests=0; roistat_is_save_data_in_cookie=1; BITRIX_CONVERSION_CONTEXT_s1=%7B%22ID%22%3A59%2C%22EXPIRE%22%3A1770843540%2C%22UNIQUE%22%3A%5B%22conversion_visit_day%22%5D%7D; _ym_isad=1; popmechanic_sbjs_migrations=popmechanic_1418474375998%3D1%7C%7C%7C1471519752600%3D1%7C%7C%7C1471519752605%3D1; _ga=GA1.1.1215700230.1770832267; _ym_visorc=w; searchbooster_v2_user_id=wgqUQzPEUYPJt5Bj8j8hz_L9pNuaAqzxTW76tu6xwlo%7C1.11.20.51; ageCheckPopupRedirectUrl=%2Fv2-mount-input; ___dc=b6614b5d-e35c-4591-ac6c-2804698aee3a; roistat_call_tracking=1; roistat_emailtracking_email=null; roistat_emailtracking_tracking_email=null; roistat_emailtracking_emails=null; roistat_cookies_to_resave=roistat_ab%2Croistat_ab_submit%2Croistat_visit%2Croistat_call_tracking%2Croistat_emailtracking_email%2Croistat_emailtracking_tracking_email%2Croistat_emailtracking_emails; BITRIX_SM_LAST_VISIT=11.02.2026%2021%3A11%3A38; qrator_jsid2=v2.0.1770832263.970.57ff104fOgiHOQCL|PQLDJTBDLkxYfKmu|IJQ93rbc/OmKxgSKa8puZw6wxj+ZRJdYbONkEPPAAdS64lT2xeMsihWehBNE+/a2V9qFLj54SU/0hTaYHbU94cYDmCbweP0dh9euUpthKAzH4iBE9PfwtpmVLDxIYMVJNsvqehQ1UbWFLa95gsUGnOiVgKt5c5TgNIDTRwfMSEw=-cjUjEJkds8d2dm40mMf6pjaldRU=; _gr_session=%7B%22s_id%22%3A%22eeba3b17-e14b-445f-a662-a1bf9c772a50%22%2C%22s_time%22%3A1770833500750%7D; _ga_7FDZ6RES9C=GS2.1.s1770832266$o1$g1$t1770833525$j34$l0$h0")
	})

	var categories []category
	c.OnHTML("div[class=newheader-catalog__wrap] > ul > li > a", func(e *colly.HTMLElement) {

		link, ok := e.DOM.Attr("href")
		if !ok {
			return
		}

		assembledCategory := category{
			link:               URL + link,
			name:               strings.TrimSpace(e.DOM.Text()),
			parentCategoryName: strings.TrimSpace(e.DOM.Parent().Parent().Parent().Prev().Text()),
		}

		categories = append(categories, assembledCategory)

	})

	if err := c.Visit("file://./stomatorg/cmd/index.html"); err != nil {
		return nil, err
	}
	return categories, nil
}
