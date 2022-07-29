import json
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
        mission_id = getattr(self, "mission_id", None)
        item = TSPersonalStatItem()
        item['mission_id'] = mission_id
        item['nick'] = nick
        item['source'] = 'thunder_skill'
        item['updated_at'] = datetime.utcnow().strftime('%Y-%m-%d %H:%M:%S')
        item['http_status'] = response.status
        if response.status == 200:
            json_body = response.json()
            content = json.dumps(json_body['stats'], ensure_ascii=False)
            item['content'] = content
            item['found'] = True
        else:
            item['content'] = response.text
            item['found'] = False
        yield item
