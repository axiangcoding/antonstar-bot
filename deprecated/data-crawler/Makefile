.PHONY:run-ts-person,run-gaijin-person,run-gaijin-cloudflare-person,gen-new-spidre

run-ts-person:
	scrapy runspider .\crawler\spiders\thunderskill.py -a nick=${nick}

run-gaijin-person:
	scrapy runspider .\crawler\spiders\gaijin.py -a nick=${nick}

run-gaijin-cloudflare-person:
	scrapy runspider .\crawler\spiders\gaijin_cloudflare.py -a nick=${nick}

gen-new-spider:
	scrapy genspider ${spider} ${domain}





