from datetime import datetime

import scrapy

from crawler.items import GaijinPersonalItem


class GaijinSpider(scrapy.Spider):
    name = 'gaijin'
    allowed_domains = ['gaijin.net']

    def start_requests(self):
        nick = getattr(self, "nick", None)
        if nick is not None:
            url = f'https://warthunder.com/zh/community/userinfo/?nick={nick}'
            yield scrapy.Request(url, method='POST')

    def parse(self, response):
        nick = getattr(self, "nick", None)
        query_id = getattr(self, "query_id", None)
        item = GaijinPersonalItem()
        item['query_id'] = query_id
        item['nick'] = nick
        item['source'] = 'gaijin'
        item['created_at'] = datetime.utcnow().strftime('%Y-%m-%d %H:%M:%S')
        item['http_status'] = response.status
        item['content'] = response.xpath("//div[@class='user-info']").extract_first()
        yield item
