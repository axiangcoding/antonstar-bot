import time

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
        print(response.request.headers)
        item = GaijinPersonalItem()
        item['createAt'] = int(time.time())
        item['http_status'] = response.status
        item['http_response'] = response.xpath("//div[@class='user-info']").extract()
        print(item['http_response'])
        yield item
