package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"runtime"
	"regexp"
	"strings"
	"github.com/kennygrant/sanitize"
	"crypto/tls"
)


type Response struct {
	url      string
	body     string
	index    int
}

func MaxParallelism() int {
	maximumProcessors := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maximumProcessors < numCPU {
		return maximumProcessors
	}
	return numCPU
}

func processResponse (resp *http.Response, url string, styleTagRegex *regexp.Regexp, emptyBrackets *regexp.Regexp, trailingTag *regexp.Regexp, pound *regexp.Regexp, emptyPtag *regexp.Regexp) (string, error) {

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		return "", err
	}

	price := doc.Find("span#priceblock_ourprice").Text()
	asin, _ := doc.Find("input#ASIN").Attr("value")
	productTitle := doc.Find("#productTitle").Text()
	customerAverageRating := doc.Find("#avgRating").Text()

	categoriesFragment := make(map[int]string)

	doc.Find("li.breadcrumb").Each(func(i int, s *goquery.Selection) {

		categoriesFragment[i] = s.Text()

	})

	categories := make([]string, len(categoriesFragment), len(categoriesFragment))

	for  i, value := range categoriesFragment {

		categories[i] = value

	}

	salesRank, _ := doc.Find("#SalesRank").Html()

	salesRank = styleTagRegex.ReplaceAllString(salesRank, "")

	salesRank = strings.Replace(salesRank, "Amazon Best Sellers Rank:", "", -1)

	salesRank = sanitize.HTML(salesRank)

	salesRank = emptyBrackets.ReplaceAllString(salesRank, "")

	salesRank = trailingTag.ReplaceAllString(salesRank, "&nbsp;&rtrif;&nbsp;")

	salesRank = pound.ReplaceAllString(salesRank, "</p><p>#")

	salesRank = "<p>" + salesRank

	salesRank = strings.Replace(salesRank, "<p></p>", "", -1)

	salesRank = salesRank + "</p>"

	salesRank = emptyPtag.ReplaceAllString(salesRank, "")

	//url := ""

	out := "<table class='table table-bordered'>" +
			"<tbody>" +
			"<tr>" +
			"<th style='width: 15%'>ASIN</th>" +
			"<td class='asin'>" + asin + "</td>" +
			"</tr>" +
			"<tr>" +
			"<th style='width: 15%'>Product Name</th>" +
			"<td>" + productTitle + "</td>" +
			"</tr>" +
			"<tr>" +
			"<th style='width: 15%'>Category</th>" +
			"<td>" + strings.Join(categories, "&nbsp;&rtrif;&nbsp;") + "</td>" +
			"</tr>" +
			"<tr>" +
			"<th style='width: 15%'>Price</th>" +
			"<td class='price'>" + price + "</td>" +
			"</tr>" +
			"<tr>" +
			"<th style='width: 15%'>Product Rating</th>" +
			"<td>" + salesRank + "</td>" +
			"</tr>" +
			"<tr>" +
			"<th style='width: 15%'>Customer Avg Rating</th>" +
			"<td>" + customerAverageRating + "</td>" +
			"</tr>" +
			"<tr>" +
			"<th style='width: 15%'>Fetched URL</th>" +
			"<td><a href='"+ url +"'>" + url + "</a></td>" +
			"</tr>" +
			"</tbody>" +
			"</table>" +
			"<hr />";

	return out, err
}

func getTemplate(client *http.Client) (string, error) {

	templateResponse, err := client.Get("https://raw.githubusercontent.com/nicholasnet/crawler/master/output.html")

	if err != nil {

		return "dd", err

	} else {

		defer templateResponse.Body.Close()
		templateContents, err := ioutil.ReadAll(templateResponse.Body)

		if err != nil {

			fmt.Printf("%s", err)
			return "dd", err

		} else {

			finalOutput := string(templateContents)
			return finalOutput, nil

		}
	}
}

func main() {

	runtime.GOMAXPROCS(MaxParallelism())

	urls := []string{
		"http://www.amazon.com/gp/product/B00NIYJF6U/ref=s9_ri_gw_g421_i1?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-5&pf_rd_r=1N9KTQ7FGVP4DBK5BE03&pf_rd_t=36701&pf_rd_p=1970555782&pf_rd_i=desktop",
		"http://www.amazon.com/SanDisk-Extreme-Memory-Adapter-SDSDQXN-032G-G46A-Version/dp/B00M55BS8G/ref=pd_sim_p_2?ie=UTF8&refRID=134NTHA1V8W0F2E0A6M7",
		"http://www.amazon.com/dp/B00KQ5A7E8?psc=1",
		"http://www.amazon.com/dp/B00DYQQSSK?psc=1",
		"http://www.amazon.com/Fitbit-Charge-Wireless-Activity-Wristband/dp/B00N2BVOUE/ref=sr_1_1?s=electronics&ie=UTF8&qid=1421723642&sr=1-1&keywords=fitbit",
		"http://www.amazon.com/Fitbit-Wireless-Activity-Tracker-Charcoal/dp/B0095PZHZE/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723642&sr=1-4&keywords=fitbit",
		"http://www.amazon.com/Sony-LT30at-Unlocked-Android-Smartphone/dp/B00L4KYKDS/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723742&sr=1-2&keywords=sony",
		"http://www.amazon.com/Beats-Studio-Wired-Over-Ear-Headphones/dp/B00E9262IE/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723766&sr=1-2&keywords=beats",
		"http://www.amazon.com/Beats-urBeats-In-Ear-Headphones-White/dp/B008CQVSXC/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723766&sr=1-4&keywords=beats",
		"http://www.amazon.com/gp/product/B00GQB1JES/ref=s9_simh_gw_p364_d0_i6?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-4&pf_rd_r=1EDQW9M33ZWHJQ6C689G&pf_rd_t=36701&pf_rd_p=1970566762&pf_rd_i=desktop",
		"http://www.amazon.com/Apple-iPhone-Space-Gray-Unlocked/dp/B00NQGP42Y/ref=sr_1_1?s=wireless&ie=UTF8&qid=1421724044&sr=1-1&keywords=iphone",
		"http://www.amazon.com/Advance-Unlocked-Dual-Phone-Black/dp/B00GXHPN1U/ref=lp_2407749011_1_3?s=wireless&ie=UTF8&qid=1421724072&sr=1-3",
		"http://www.amazon.com/dp/B00M6TLHTQ?psc=1",
		"http://www.amazon.com/LG-Realm-LS620-Contract-Mobile/dp/B00N15E6TW/ref=acs_ux_rw_ts_e_2407748011_2?ie=UTF8&s=electronics&pf_rd_p=1964575062&pf_rd_s=merchandised-search-7&pf_rd_t=101&pf_rd_i=2407748011&pf_rd_m=ATVPDKIKX0DER&pf_rd_r=1ZF0JRZH5T4AN9P8WZ4P",
		"http://www.amazon.com/LG-Volt-Prepaid-Phone-Mobile/dp/B00K8CS8VS/ref=pd_sim_e_7?ie=UTF8&refRID=0FQVAMZTZFCR6SJHZ276",
		"http://www.amazon.com/AERO-ARMOR-Protective-Case-LS740/dp/B00KLS2982/ref=pd_sim_cps_7?ie=UTF8&refRID=0YBX5VQG5XAZ3RFMSNA3",
		"http://www.amazon.com/Skinomi%C2%AE-TechSkin-Replacement-Definition-Anti-Bubble/dp/B00IT75WPY/ref=pd_sim_cps_8?ie=UTF8&refRID=0VB9RV2A1TSGBM4HB9H8",
		"http://www.amazon.com/Sunny-Health-Fitness-Mini-Cycle/dp/B0016BQFV0/ref=sr_1_9?ie=UTF8&qid=1421724292&sr=8-9&keywords=cycle",
		"http://www.amazon.com/Drive-Medical-Exerciser-Attractive-Silver/dp/B002VWK09Q/ref=pd_sim_sg_8?ie=UTF8&refRID=04GGD3YDBN60WHWDQY8C",
		"http://www.amazon.com/dp/B00NA91ENU?psc=1",
		"http://www.amazon.com/gp/product/B00NIYJF6U/ref=s9_ri_gw_g421_i1?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-5&pf_rd_r=1N9KTQ7FGVP4DBK5BE03&pf_rd_t=36701&pf_rd_p=1970555782&pf_rd_i=desktop",
		"http://www.amazon.com/SanDisk-Extreme-Memory-Adapter-SDSDQXN-032G-G46A-Version/dp/B00M55BS8G/ref=pd_sim_p_2?ie=UTF8&refRID=134NTHA1V8W0F2E0A6M7",
		"http://www.amazon.com/dp/B00KQ5A7E8?psc=1",
		"http://www.amazon.com/dp/B00DYQQSSK?psc=1",
		"http://www.amazon.com/Fitbit-Charge-Wireless-Activity-Wristband/dp/B00N2BVOUE/ref=sr_1_1?s=electronics&ie=UTF8&qid=1421723642&sr=1-1&keywords=fitbit",
		"http://www.amazon.com/Fitbit-Wireless-Activity-Tracker-Charcoal/dp/B0095PZHZE/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723642&sr=1-4&keywords=fitbit",
		"http://www.amazon.com/Sony-LT30at-Unlocked-Android-Smartphone/dp/B00L4KYKDS/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723742&sr=1-2&keywords=sony",
		"http://www.amazon.com/Beats-Studio-Wired-Over-Ear-Headphones/dp/B00E9262IE/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723766&sr=1-2&keywords=beats",
		"http://www.amazon.com/Beats-urBeats-In-Ear-Headphones-White/dp/B008CQVSXC/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723766&sr=1-4&keywords=beats",
		"http://www.amazon.com/gp/product/B00GQB1JES/ref=s9_simh_gw_p364_d0_i6?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-4&pf_rd_r=1EDQW9M33ZWHJQ6C689G&pf_rd_t=36701&pf_rd_p=1970566762&pf_rd_i=desktop",
		"http://www.amazon.com/Apple-iPhone-Space-Gray-Unlocked/dp/B00NQGP42Y/ref=sr_1_1?s=wireless&ie=UTF8&qid=1421724044&sr=1-1&keywords=iphone",
		"http://www.amazon.com/Advance-Unlocked-Dual-Phone-Black/dp/B00GXHPN1U/ref=lp_2407749011_1_3?s=wireless&ie=UTF8&qid=1421724072&sr=1-3",
		"http://www.amazon.com/dp/B00M6TLHTQ?psc=1",
		"http://www.amazon.com/LG-Realm-LS620-Contract-Mobile/dp/B00N15E6TW/ref=acs_ux_rw_ts_e_2407748011_2?ie=UTF8&s=electronics&pf_rd_p=1964575062&pf_rd_s=merchandised-search-7&pf_rd_t=101&pf_rd_i=2407748011&pf_rd_m=ATVPDKIKX0DER&pf_rd_r=1ZF0JRZH5T4AN9P8WZ4P",
		"http://www.amazon.com/LG-Volt-Prepaid-Phone-Mobile/dp/B00K8CS8VS/ref=pd_sim_e_7?ie=UTF8&refRID=0FQVAMZTZFCR6SJHZ276",
		"http://www.amazon.com/AERO-ARMOR-Protective-Case-LS740/dp/B00KLS2982/ref=pd_sim_cps_7?ie=UTF8&refRID=0YBX5VQG5XAZ3RFMSNA3",
		"http://www.amazon.com/Skinomi%C2%AE-TechSkin-Replacement-Definition-Anti-Bubble/dp/B00IT75WPY/ref=pd_sim_cps_8?ie=UTF8&refRID=0VB9RV2A1TSGBM4HB9H8",
		"http://www.amazon.com/Sunny-Health-Fitness-Mini-Cycle/dp/B0016BQFV0/ref=sr_1_9?ie=UTF8&qid=1421724292&sr=8-9&keywords=cycle",
		"http://www.amazon.com/Drive-Medical-Exerciser-Attractive-Silver/dp/B002VWK09Q/ref=pd_sim_sg_8?ie=UTF8&refRID=04GGD3YDBN60WHWDQY8C",
		"http://www.amazon.com/dp/B00NA91ENU?psc=1",
		"http://www.amazon.com/gp/product/B00NIYJF6U/ref=s9_ri_gw_g421_i1?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-5&pf_rd_r=1N9KTQ7FGVP4DBK5BE03&pf_rd_t=36701&pf_rd_p=1970555782&pf_rd_i=desktop",
		"http://www.amazon.com/SanDisk-Extreme-Memory-Adapter-SDSDQXN-032G-G46A-Version/dp/B00M55BS8G/ref=pd_sim_p_2?ie=UTF8&refRID=134NTHA1V8W0F2E0A6M7",
		"http://www.amazon.com/dp/B00KQ5A7E8?psc=1",
		"http://www.amazon.com/dp/B00DYQQSSK?psc=1",
		"http://www.amazon.com/Fitbit-Charge-Wireless-Activity-Wristband/dp/B00N2BVOUE/ref=sr_1_1?s=electronics&ie=UTF8&qid=1421723642&sr=1-1&keywords=fitbit",
		"http://www.amazon.com/Fitbit-Wireless-Activity-Tracker-Charcoal/dp/B0095PZHZE/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723642&sr=1-4&keywords=fitbit",
		"http://www.amazon.com/Sony-LT30at-Unlocked-Android-Smartphone/dp/B00L4KYKDS/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723742&sr=1-2&keywords=sony",
		"http://www.amazon.com/Beats-Studio-Wired-Over-Ear-Headphones/dp/B00E9262IE/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723766&sr=1-2&keywords=beats",
//		"http://www.amazon.com/Beats-urBeats-In-Ear-Headphones-White/dp/B008CQVSXC/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723766&sr=1-4&keywords=beats",
//		"http://www.amazon.com/gp/product/B00GQB1JES/ref=s9_simh_gw_p364_d0_i6?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-4&pf_rd_r=1EDQW9M33ZWHJQ6C689G&pf_rd_t=36701&pf_rd_p=1970566762&pf_rd_i=desktop",
//		"http://www.amazon.com/Apple-iPhone-Space-Gray-Unlocked/dp/B00NQGP42Y/ref=sr_1_1?s=wireless&ie=UTF8&qid=1421724044&sr=1-1&keywords=iphone",
	}

	output := make(chan Response, len(urls))

	var wg sync.WaitGroup

	wg.Add(len(urls) + 1)

	counter := 0

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	styleTagRegex := regexp.MustCompile("<style type=\"text\\/css\">(.|\\n)*<\\/style>")
	emptyBrackets := regexp.MustCompile("\\(.*\\)")
	trailingTag := regexp.MustCompile("\\>")
	pound := regexp.MustCompile("#")
	emptyPtag := regexp.MustCompile("<p>\\s*</p>")


	for _, url := range urls {
		go func(url string, counter *int, client *http.Client) {

			defer wg.Done()

			req, err := http.NewRequest("GET", url, nil)

			if err != nil {

				log.Fatal(err)
				output <- Response{url, "", *counter}

			} else {

				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0")
				resp, err := client.Do(req)

				if err != nil {

					log.Fatal(err)
					output <- Response{url, "", *counter}

				} else {

					out, err := processResponse(resp, url, styleTagRegex, emptyBrackets, trailingTag, pound, emptyPtag)

					defer resp.Body.Close()

					if err != nil {

						log.Fatal(err)
						output <- Response{url, "", *counter}

					} else {

						output <- Response{url, out, *counter}
					}

				}

			}

			*counter = *counter + 1

		}(url, &counter, client)

	}

	go func(client *http.Client) {

		itemsToScrape := len(urls) - 1

		start := time.Now()

		mergedResult := ""

		for response := range output {

			mergedResult = mergedResult + response.body

			if response.index == itemsToScrape {

				elapsed := time.Since(start)
				log.Printf("Scrapping took %s", elapsed)

				template, _ := getTemplate(client)
				finalOutput := strings.Replace(template, "<!-- OUTPUT -->", mergedResult, -1)
				_ = ioutil.WriteFile("output.html", []byte(finalOutput), 0644)

				wg.Done()
			}
		}

	}(client)

	wg.Wait()
}
