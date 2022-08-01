import json
from datetime import datetime

import scrapy

from crawler.items import GaijinPersonalItem
from util.extract_gaijin import extract_user_info


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
        mission_id = getattr(self, "mission_id", None)
        item = GaijinPersonalItem()
        item['mission_id'] = mission_id
        item['nick'] = nick
        item['source'] = 'gaijin'
        item['updated_at'] = datetime.utcnow().strftime('%Y-%m-%d %H:%M:%S')
        item['http_status'] = response.status

        user_not_found = response.xpath("//div[@class='user__unavailable-title']").extract()
        if len(user_not_found) > 0:
            item['content'] = 'not found'
            item['found'] = False
        else:
            item['found'] = True
            content = extract_user_info(response.body)
            item['content'] = json.dumps(content, ensure_ascii=False)
        yield item

