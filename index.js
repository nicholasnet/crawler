console.time("Script took");

var request = require("request"),
    cheerio = require("cheerio"),
    fs = require("fs"),
    urls = [
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
        'http://www.amazon.com/Beats-Studio-Wired-Over-Ear-Headphones/dp/B00E9262IE/ref=sr_1_2?s=electronics&ie=UTF8&qid=1421723766&sr=1-2&keywords=beats'/*,
        'http://www.amazon.com/Beats-urBeats-In-Ear-Headphones-White/dp/B008CQVSXC/ref=sr_1_4?s=electronics&ie=UTF8&qid=1421723766&sr=1-4&keywords=beats',
        'http://www.amazon.com/gp/product/B00GQB1JES/ref=s9_simh_gw_p364_d0_i6?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-4&pf_rd_r=1EDQW9M33ZWHJQ6C689G&pf_rd_t=36701&pf_rd_p=1970566762&pf_rd_i=desktop',
        'http://www.amazon.com/Apple-iPhone-Space-Gray-Unlocked/dp/B00NQGP42Y/ref=sr_1_1?s=wireless&ie=UTF8&qid=1421724044&sr=1-1&keywords=iphone'*/
    ],
    urlsCount = urls.length,
    output = '',
    counter = 0,
    baseRequest = request.defaults({
        headers: {
            "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
            "Accept-Language": "en-US,en;q=0.5"
        }
    }),
    errorCount = 0;

urls.forEach(function(item, index) {

    baseRequest(item, function (error, response, body) {

        if (!error && response.statusCode == 200) {
            $ = cheerio.load(body);
            var asin = $('input#ASIN').val();
            //var asin = $('#fbt_x_title').data('asin'); // without useragent
            var price = $('#priceblock_ourprice').text();
            //var price = $('#actualPriceValue').text(); // without useragent
            var productTitle = $('#productTitle').text();
            //var productTitle = $('#btAsinTitle').text(); // without useragent
            var category = [];

            $('li.breadcrumb').each(function(i, elem) {
                category[i] = $(this).text();
            });
            category = category.join('&nbsp;&rtrif;&nbsp;');

            var salesRank = $('#SalesRank');

            if (salesRank.length !== 0) {
                salesRank = salesRank.find('style').remove().end().text();
                salesRank = salesRank.replace('Amazon Best Sellers Rank:','');
                salesRank = salesRank.replace(/\(.*\)/,'');
                salesRank = salesRank.replace(/\>/g,'&nbsp;&rtrif;&nbsp;');
                salesRank = salesRank.replace(/#/g,'</p><p>#');
                salesRank = '<p>' + salesRank;
                salesRank = salesRank.replace('<p></p>','');
                salesRank = salesRank + '</p>';
                salesRank = salesRank.replace(/<p>\s*<\/p>/g,'');
            }


            var customerAverageRating = $('#avgRating').text();
            //var customerAverageRating = $('div.acrRating').text();

            var fragment = '<table class="table table-bordered">' +
                               '<tbody>' +
                                   '<tr>' +
                                       '<th style="width: 15%">ASIN</th>' +
                                       '<td class="asin">' + asin + '</td>' +
                                   '</tr>' +
                                   '<tr>' +
                                       '<th style="width: 15%">Product Name</th>' +
                                       '<td>' + productTitle + '</td>' +
                                   '</tr>' +
                                   '<tr>' +
                                       '<th style="width: 15%">Category</th>' +
                                       '<td>' + category + '</td>' +
                                   '</tr>' +
                                   '<tr>' +
                                       '<th style="width: 15%">Price</th>' +
                                       '<td class="price">' + price + '</td>' +
                                   '</tr>' +
                                   '<tr>' +
                                       '<th style="width: 15%">Product Rating</th>' +
                                       '<td>' + salesRank + '</td>' +
                                   '</tr>' +
                                   '<tr>' +
                                       '<th style="width: 15%">Customer Avg Rating</th>' +
                                       '<td>' + customerAverageRating + '</td>' +
                                   '</tr>' +
                                   '<tr>' +
                                       '<th style="width: 15%">Fetched URL</th>' +
                                       '<td><a href="' + item + '">' + item + '</a></td>' +
                                   '</tr>' +
                               '</tbody>' +
                            '</table>' +
                            '<hr />';

            output = output + fragment;

        }
        else {
            errorCount++;
            console.log("Request failed");
        }

        counter++;


        if (counter === urlsCount) {
            console.log("Tried to fetch " + urls.length + ". Unable to fetch " + errorCount);
            console.timeEnd("Script took");
            baseRequest('https://raw.githubusercontent.com/nicholasnet/crawler/master/output.html', function (error, response, html) {
                var $ = cheerio.load(html);
                $('#output').html(output);

                fs.writeFile('output.html', $.html(), function (err) {
                    console.log('HTML file is saved!');
                });
            });
        }
    });
});
