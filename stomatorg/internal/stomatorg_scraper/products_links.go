package stomatorg_scraper

import (
	"fmt"
	"time"
)

func (p *Page) parseProductLinks(categories []category) ([]product, error) {
	products := make([]product, 0, 2000)
	for _, cg := range categories {
		time.Sleep(time.Second)
		parseProducts, err := p.parseProductsLinkCategory(cg)
		if err != nil {
			return nil, err
		}

		products = append(products, parseProducts...)

		break // !!!
	}
	return products, nil
}

func (p *Page) parseProductsLinkCategory(cg category) ([]product, error) {
	//c := colly.NewCollector()
	//
	//c.OnRequest(func(r *colly.Request) {
	//	r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	//	r.Headers.Set("Accept-Language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	//	r.Headers.Set("Cache-Control", "no-cache")
	//	r.Headers.Set("Connection", "keep-alive")
	//	r.Headers.Set("Pragma", "no-cache")
	//	r.Headers.Set("Referer", "https://stomatorg.ru/catalog/anesteziya/?utm_referrer=https%3A%2F%2Fstomatorg.ru%2Fcatalog%2Fstomatologicheskie_materialy%2F&roistat_referrer=https://stomatorg.ru/catalog/stomatologicheskie_materialy/")
	//	r.Headers.Set("Sec-Fetch-Dest", "document")
	//	r.Headers.Set("Sec-Fetch-Mode", "navigate")
	//	r.Headers.Set("Sec-Fetch-Site", "same-origin")
	//	r.Headers.Set("Sec-Fetch-User", "?1")
	//	r.Headers.Set("Upgrade-Insecure-Requests", "1")
	//	r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 YaBrowser/25.12.0.0 Safari/537.36")
	//	r.Headers.Set("sec-ch-ua", "\"Chromium\";v=\"142\", \"YaBrowser\";v=\"25.12\", \"Not_A Brand\";v=\"99\", \"Yowser\";v=\"2.5\"")
	//	r.Headers.Set("sec-ch-ua-mobile", "?0")
	//	r.Headers.Set("sec-ch-ua-platform", "\"Linux\"")
	//	r.Headers.Set("Cookie", "\nBITRIX_SM_GUEST_ID=34065829; mindboxDeviceUUID=c26a2121-46f6-412e-aa7e-cdcba7acba2e; directCrm-session=%7B%22deviceGuid%22%3A%22c26a2121-46f6-412e-aa7e-cdcba7acba2e%22%7D; _ym_uid=1770832266394185598; _ym_d=1770832266; roistat_visit=7247016; roistat_first_visit=7247016; roistat_visit_cookie_expire=1209600; roistat_is_need_listen_requests=0; roistat_is_save_data_in_cookie=1; BITRIX_CONVERSION_CONTEXT_s1=%7B%22ID%22%3A59%2C%22EXPIRE%22%3A1770843540%2C%22UNIQUE%22%3A%5B%22conversion_visit_day%22%5D%7D; _ym_isad=1; popmechanic_sbjs_migrations=popmechanic_1418474375998%3D1%7C%7C%7C1471519752600%3D1%7C%7C%7C1471519752605%3D1; _ga=GA1.1.1215700230.1770832267; searchbooster_v2_user_id=wgqUQzPEUYPJt5Bj8j8hz_L9pNuaAqzxTW76tu6xwlo%7C1.11.20.51; ageCheckPopupRedirectUrl=%2Fv2-mount-input; ___dc=b6614b5d-e35c-4591-ac6c-2804698aee3a; roistat_call_tracking=1; roistat_emailtracking_email=null; roistat_emailtracking_tracking_email=null; roistat_emailtracking_emails=null; roistat_cookies_to_resave=roistat_ab%2Croistat_ab_submit%2Croistat_visit%2Croistat_call_tracking%2Croistat_emailtracking_email%2Croistat_emailtracking_tracking_email%2Croistat_emailtracking_emails; PHPSESSID=8Orw4ppHHHtL5h9Q1O5fHFo50nVMFbTV; _ym_visorc=w; ESR_REDIRECTS=YTowOnt9; qrator_jsid2=v2.0.1770835885.990.57ff104fSYtw5ver|4otiT5NnZLbUWcAz|0ejxtNpJU1oeGK+YPB6BJowfd0Mdd+gT8h6vWZxsO8xX7TZnTuLynbERUhgjqpTmjNnmPw/6kd9JnpPQeKjiRg49PphPrl9t+zHI7ZKTgIYBR7fPbTAcg9b752Z6vxWCLOeEMkLl37kHzT5Aw5z5CsmdL1NDaVhF3uyhsOKzs3k=-tsmFpABVzVmYtee6pV9Wta1eIWI=; ifCookieAgree=Y; BITRIX_SM_LAST_VISIT=11.02.2026%2022%3A17%3A37; _gr_session=%7B%22s_id%22%3A%22eeba3b17-e14b-445f-a662-a1bf9c772a50%22%2C%22s_time%22%3A1770837459050%7D; _ga_7FDZ6RES9C=GS2.1.s1770835887$o2$g1$t1770837459$j37$l0$h0")
	//
	//})
	//
	//var products []product
	//c.OnHTML("div[class=container]", func(e *colly.HTMLElement) {
	//	fmt.Println("$+$+$+$")
	//
	//	link, ok := e.DOM.Find("a[class=catalog-elements-item__name]").Attr("href")
	//	if !ok {
	//		return
	//	}
	//
	//	assembledProduct := product{
	//		link: URL + link,
	//	}
	//
	//	products = append(products, assembledProduct)
	//
	//	nextPageLink, ok := e.DOM.Find("a[class=navigation_1_next_page]").Attr("href")
	//	if !ok {
	//		return
	//	}
	//
	//	if err := e.Request.Visit(nextPageLink); err != nil {
	//		return
	//	}
	//})
	//
	//fmt.Println(cg.link)
	//if err := c.Visit(cg.link); err != nil {
	//	return nil, err
	//}

	var products []product
	visitLink := cg.link
	for {

		if _, err := p.Goto("https://viewdns.info/iphistory/?domain=stomatorg.ru"); err != nil {
			return nil, fmt.Errorf("could not goto: %v", err)
		}
		time.Sleep(20 * time.Second)

		if _, err := p.Goto(visitLink); err != nil {
			return nil, fmt.Errorf("could not goto: %v", err)
		}

		container, err := p.Locator("div[class=container]").All()
		if err != nil {
			return nil, fmt.Errorf("could not get entries: %v", err)
		}
		for _, entry := range container {
			link, err := entry.Locator("a[class=catalog-elements-item__name]").GetAttribute("href")
			if err != nil {
				return nil, fmt.Errorf("could not get text content: %v", err)
			}

			products = append(products, product{
				category: cg,
				link:     link,
			})
		}

		nextVisitLink, err := p.Locator("a[class=navigation_1_next_page]").GetAttribute("href")
		if err != nil {
			break
			//return nil, fmt.Errorf("could not get entries: %v", err)
		}
		visitLink = nextVisitLink
	}

	return products, nil
}
