.PHONY:run-ts-person,run-gaijin-person,gen-new-spidre

run-ts-person:
	scrapy runspider .\crawler\spiders\thunderskill.py -a nick=${nick} -o thunderskill.json

run-gaijin-person:
	scrapy runspider .\crawler\spiders\gaijin.py -a nick=${nick} -o gaijin.json

gen-new-spider:
	scrapy genspider ${spider} ${domain}





