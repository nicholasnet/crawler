__author__ = 'nicholasnet'

import re
import grequests
import time

from pyquery import PyQuery

start_time = time.time()

urls = ['http://www.amazon.com/gp/product/B00NIYJF6U/ref=s9_ri_gw_g421_i1?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-5&pf_rd_r=1N9KTQ7FGVP4DBK5BE03&pf_rd_t=36701&pf_rd_p=1970555782&pf_rd_i=desktop',
        'http://www.amazon.com/SanDisk-Extreme-Memory-Adapter-SDSDQXN-032G-G46A-Version/dp/B00M55BS8G/ref=pd_sim_p_2?ie=UTF8&refRID=134NTHA1V8W0F2E0A6M7',
        'http://www.amazon.com/dp/B00KQ5A7E8?psc=1',
        'http://www.amazon.com/dp/B00DYQQSSK?psc=1',
        'http://www.amazon.com/Fitbit-Charge-Wireless-Activity-Wristband/dp/B00N2BVOUE/ref=sr_1_1?s=electronics&ie=UTF8&qid=1421723642&sr=1-1&keywords=fitbit',
        'http://www.amazon.com/Fitbit-Wireless-Activity-Tracker-Charcoal/dp/B0095PZHZE/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723642&sr=1-4&keywords=fitbit',
        'http://www.amazon.com/Sony-LT30at-Unlocked-Android-Smartphone/dp/B00L4KYKDS/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723742&sr=1-2&keywords=sony',
        'http://www.amazon.com/Beats-Studio-Wired-Over-Ear-Headphones/dp/B00E9262IE/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723766&sr=1-2&keywords=beats',
        'http://www.amazon.com/Beats-urBeats-In-Ear-Headphones-White/dp/B008CQVSXC/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723766&sr=1-4&keywords=beats',
        'http://www.amazon.com/gp/product/B00GQB1JES/ref=s9_simh_gw_p364_d0_i6?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-4&pf_rd_r=1EDQW9M33ZWHJQ6C689G&pf_rd_t=36701&pf_rd_p=1970566762&pf_rd_i=desktop',
        'http://www.amazon.com/Apple-iPhone-Space-Gray-Unlocked/dp/B00NQGP42Y/ref=sr_1_1?s=wireless&ie=UTF8&qid=1421724044&sr=1-1&keywords=iphone',
        'http://www.amazon.com/Advance-Unlocked-Dual-Phone-Black/dp/B00GXHPN1U/ref=lp_2407749011_1_3?s=wireless&ie=UTF8&qid=1421724072&sr=1-3',
        'http://www.amazon.com/dp/B00M6TLHTQ?psc=1',
        'http://www.amazon.com/LG-Realm-LS620-Contract-Mobile/dp/B00N15E6TW/ref=acs_ux_rw_ts_e_2407748011_2?ie=UTF8&s=electronics&pf_rd_p=1964575062&pf_rd_s=merchandised-search-7&pf_rd_t=101&pf_rd_i=2407748011&pf_rd_m=ATVPDKIKX0DER&pf_rd_r=1ZF0JRZH5T4AN9P8WZ4P',
        'http://www.amazon.com/LG-Volt-Prepaid-Phone-Mobile/dp/B00K8CS8VS/ref=pd_sim_e_7?ie=UTF8&refRID=0FQVAMZTZFCR6SJHZ276',
        'http://www.amazon.com/AERO-ARMOR-Protective-Case-LS740/dp/B00KLS2982/ref=pd_sim_cps_7?ie=UTF8&refRID=0YBX5VQG5XAZ3RFMSNA3',
        'http://www.amazon.com/Skinomi%C2%AE-TechSkin-Replacement-Definition-Anti-Bubble/dp/B00IT75WPY/ref=pd_sim_cps_8?ie=UTF8&refRID=0VB9RV2A1TSGBM4HB9H8',
        'http://www.amazon.com/Sunny-Health-Fitness-Mini-Cycle/dp/B0016BQFV0/ref=sr_1_9?ie=UTF8&qid=1421724292&sr=8-9&keywords=cycle',
        'http://www.amazon.com/Drive-Medical-Exerciser-Attractive-Silver/dp/B002VWK09Q/ref=pd_sim_sg_8?ie=UTF8&refRID=04GGD3YDBN60WHWDQY8C',
        'http://www.amazon.com/dp/B00NA91ENU?psc=1',
        'http://www.amazon.com/gp/product/B00NIYJF6U/ref=s9_ri_gw_g421_i1?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-5&pf_rd_r=1N9KTQ7FGVP4DBK5BE03&pf_rd_t=36701&pf_rd_p=1970555782&pf_rd_i=desktop',
        'http://www.amazon.com/SanDisk-Extreme-Memory-Adapter-SDSDQXN-032G-G46A-Version/dp/B00M55BS8G/ref=pd_sim_p_2?ie=UTF8&refRID=134NTHA1V8W0F2E0A6M7',
        'http://www.amazon.com/dp/B00KQ5A7E8?psc=1',
        'http://www.amazon.com/dp/B00DYQQSSK?psc=1',
        'http://www.amazon.com/Fitbit-Charge-Wireless-Activity-Wristband/dp/B00N2BVOUE/ref=sr_1_1?s=electronics&ie=UTF8&qid=1421723642&sr=1-1&keywords=fitbit',
        'http://www.amazon.com/Fitbit-Wireless-Activity-Tracker-Charcoal/dp/B0095PZHZE/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723642&sr=1-4&keywords=fitbit',
        'http://www.amazon.com/Sony-LT30at-Unlocked-Android-Smartphone/dp/B00L4KYKDS/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723742&sr=1-2&keywords=sony',
        'http://www.amazon.com/Beats-Studio-Wired-Over-Ear-Headphones/dp/B00E9262IE/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723766&sr=1-2&keywords=beats',
        'http://www.amazon.com/Beats-urBeats-In-Ear-Headphones-White/dp/B008CQVSXC/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723766&sr=1-4&keywords=beats',
        'http://www.amazon.com/gp/product/B00GQB1JES/ref=s9_simh_gw_p364_d0_i6?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-4&pf_rd_r=1EDQW9M33ZWHJQ6C689G&pf_rd_t=36701&pf_rd_p=1970566762&pf_rd_i=desktop',
        'http://www.amazon.com/Apple-iPhone-Space-Gray-Unlocked/dp/B00NQGP42Y/ref=sr_1_1?s=wireless&ie=UTF8&qid=1421724044&sr=1-1&keywords=iphone',
        'http://www.amazon.com/Advance-Unlocked-Dual-Phone-Black/dp/B00GXHPN1U/ref=lp_2407749011_1_3?s=wireless&ie=UTF8&qid=1421724072&sr=1-3',
        'http://www.amazon.com/dp/B00M6TLHTQ?psc=1',
        'http://www.amazon.com/LG-Realm-LS620-Contract-Mobile/dp/B00N15E6TW/ref=acs_ux_rw_ts_e_2407748011_2?ie=UTF8&s=electronics&pf_rd_p=1964575062&pf_rd_s=merchandised-search-7&pf_rd_t=101&pf_rd_i=2407748011&pf_rd_m=ATVPDKIKX0DER&pf_rd_r=1ZF0JRZH5T4AN9P8WZ4P',
        'http://www.amazon.com/LG-Volt-Prepaid-Phone-Mobile/dp/B00K8CS8VS/ref=pd_sim_e_7?ie=UTF8&refRID=0FQVAMZTZFCR6SJHZ276',
        'http://www.amazon.com/AERO-ARMOR-Protective-Case-LS740/dp/B00KLS2982/ref=pd_sim_cps_7?ie=UTF8&refRID=0YBX5VQG5XAZ3RFMSNA3',
        'http://www.amazon.com/Skinomi%C2%AE-TechSkin-Replacement-Definition-Anti-Bubble/dp/B00IT75WPY/ref=pd_sim_cps_8?ie=UTF8&refRID=0VB9RV2A1TSGBM4HB9H8',
        'http://www.amazon.com/Sunny-Health-Fitness-Mini-Cycle/dp/B0016BQFV0/ref=sr_1_9?ie=UTF8&qid=1421724292&sr=8-9&keywords=cycle',
        'http://www.amazon.com/Drive-Medical-Exerciser-Attractive-Silver/dp/B002VWK09Q/ref=pd_sim_sg_8?ie=UTF8&refRID=04GGD3YDBN60WHWDQY8C',
        'http://www.amazon.com/dp/B00NA91ENU?psc=1',
        'http://www.amazon.com/gp/product/B00NIYJF6U/ref=s9_ri_gw_g421_i1?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-5&pf_rd_r=1N9KTQ7FGVP4DBK5BE03&pf_rd_t=36701&pf_rd_p=1970555782&pf_rd_i=desktop',
        'http://www.amazon.com/SanDisk-Extreme-Memory-Adapter-SDSDQXN-032G-G46A-Version/dp/B00M55BS8G/ref=pd_sim_p_2?ie=UTF8&refRID=134NTHA1V8W0F2E0A6M7',
        'http://www.amazon.com/dp/B00KQ5A7E8?psc=1',
        'http://www.amazon.com/dp/B00DYQQSSK?psc=1',
        'http://www.amazon.com/Fitbit-Charge-Wireless-Activity-Wristband/dp/B00N2BVOUE/ref=sr_1_1?s=electronics&ie=UTF8&qid=1421723642&sr=1-1&keywords=fitbit',
        'http://www.amazon.com/Fitbit-Wireless-Activity-Tracker-Charcoal/dp/B0095PZHZE/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723642&sr=1-4&keywords=fitbit',
        'http://www.amazon.com/Sony-LT30at-Unlocked-Android-Smartphone/dp/B00L4KYKDS/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723742&sr=1-2&keywords=sony',
        'http://www.amazon.com/Beats-Studio-Wired-Over-Ear-Headphones/dp/B00E9262IE/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723766&sr=1-2&keywords=beats']

remove_whitespace_around_brackets = re.compile("\(.*\)", re.IGNORECASE)
remove_trailing_tag = re.compile(">", re.IGNORECASE)
remove_pound_symbol = re.compile("#", re.IGNORECASE)
remove_p_tag = re.compile("<p></p>")
remove_blank_p_tag = re.compile("<p>\s*<\/p>")

template_url = 'https://raw.githubusercontent.com/nicholasnet/crawler/master/output.html'


headers = {'User-Agent':'Mozilla/5.0 (Windows NT 6.1; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0',
           'Accept':'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
           'Accept-Language':'en-US,en;q=0.5'}

requests = (grequests.get(u, headers=headers, stream=False) for u in urls)


output = ""

for response in grequests.map(requests):

    if response.raise_for_status() == None:

        item = response.url

        pq = PyQuery(response.content)

        price = pq('#priceblock_ourprice').text()

        asin = pq('input#ASIN').val()

        product_title = pq('#productTitle').text()

        customer_average_rating = pq('#avgRating').text()

        category = []

        for fragment in pq('li.breadcrumb'):

            category.append(pq(fragment).text())

        category = '&nbsp;&rtrif;&nbsp;'.join(category)

        sales_rank = pq('#SalesRank')

        if len(sales_rank) != 0:

            sales_rank_node = PyQuery(sales_rank)
            sales_rank = sales_rank_node.find('style').remove()
            sales_rank = sales_rank_node.text()
            sales_rank = sales_rank.replace('Amazon Best Sellers Rank:', '')
            sales_rank = remove_whitespace_around_brackets.sub('', sales_rank)
            sales_rank = remove_trailing_tag.sub('&nbsp;&rtrif;&nbsp;', sales_rank)
            sales_rank = remove_pound_symbol.sub('</p><p>#', sales_rank)
            sales_rank = '<p>' + sales_rank
            sales_rank = sales_rank.replace('<p></p>', '')
            sales_rank = sales_rank + '</p>'
            sales_rank = remove_blank_p_tag.sub('', sales_rank)
        else:

            sales_rank = ''

        fragment = '<table class="table table-bordered">'
        fragment += '<tbody>'
        fragment += '<tr>'
        fragment += '<th style="width: 15%">ASIN</th>'
        fragment += '<td class="asin">' + asin + '</td>'
        fragment += '</tr>'
        fragment += '<tr>'
        fragment += '<th style="width: 15%">Product Name</th>'
        fragment += '<td>' + product_title + '</td>'
        fragment += '</tr>'
        fragment += '<tr>'
        fragment += '<th style="width: 15%">Category</th>'
        fragment += '<td>' + category + '</td>'
        fragment += '</tr>'
        fragment += '<tr>'
        fragment += '<th style="width: 15%">Price</th>'
        fragment += '<td class="price">' + price + '</td>'
        fragment += '</tr>'
        fragment += '<tr>'
        fragment += '<th style="width: 15%">Product Rating</th>'
        fragment += '<td>' + sales_rank + '</td>'
        fragment += '</tr>'
        fragment += '<tr>'
        fragment += '<th style="width: 15%">Customer Avg Rating</th>'
        fragment += '<td>' + customer_average_rating + '</td>'
        fragment += '</tr>'
        fragment += '<tr>'
        fragment += '<th style="width: 15%">Fetched URL</th>'
        fragment += '<td><a href="' + item + '">' + item + '</a></td>'
        fragment += '</tr>'
        fragment += '</tbody>'
        fragment += '</table>'
        fragment += '<hr />'
        output += fragment

    response.close()

output = output.encode('utf-8').strip()

end_time = time.time()

template_request = [grequests.get(template_url, headers=headers, verify=False, stream=False)]

for template_response in grequests.map(template_request):

    if template_response.raise_for_status() != None:

        template_content = template_response.content
        template_content = template_content.encode('utf-8').strip()
        output = template_content.replace('<!-- OUTPUT -->', output)

    template_response.close()

file = open("output.html", "w")
file.write(output)
file.close()
elapsed_time = end_time - start_time

print 'Script took ' + str(elapsed_time)
