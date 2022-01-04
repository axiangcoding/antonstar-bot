import time

import scrapy

from crawler.items import TSPersonalStatItem


class ThunderSkillSpider(scrapy.Spider):
    name = 'quotes'
    allowed_domains = ['thunderskill.com']

    def start_requests(self):
        nick = getattr(self, "nick", None)
        if nick is not None:
            url = f'https://thunderskill.com/en/stat/{nick}/export/json'
            yield scrapy.Request(url)

    def parse(self, response):
        item = TSPersonalStatItem()
        item['createAt']= int(time.time())
        item['http_status'] = response.status
        item['http_response'] = response.text
        yield item
