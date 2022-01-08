from datetime import datetime

import scrapy

from crawler.items import TSPersonalStatItem


class ThunderSkillSpider(scrapy.Spider):
    name = 'thunderskill'
    allowed_domains = ['thunderskill.com']

    def start_requests(self):
        nick = getattr(self, "nick", None)
        if nick is not None:
            url = f'https://thunderskill.com/en/stat/{nick}/export/json'
            yield scrapy.Request(url)

    def parse(self, response):
        nick = getattr(self, "nick", None)
        item = TSPersonalStatItem()
        item['nick'] = nick
        item['source'] = 'thunder skill'
        item['created_at'] = datetime.utcnow().strftime('%Y-%m-%d %H:%M:%S')
        item['http_status'] = response.status
        item['content'] = response.text
        yield item
