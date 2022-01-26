import json
import time
from datetime import datetime

import scrapy

from crawler.items import GaijinPersonalItem
from util import string
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
        query_id = getattr(self, "query_id", None)
        item = GaijinPersonalItem()
        item['query_id'] = query_id
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


def extract_register_date(date_str: str):
    return time.strftime("%Y-%m-%d", time.strptime(date_str.split(" ")[-1], "%d.%m.%Y"))


def extract_level(level_str: str):
    return level_str.replace("Level", "").replace("等级", "")


def zip_user_stat(lst: list, split: int):
    title = lst[:split]
    ret_lst = []
    for i in range(split, len(lst), split):
        line = lst[i:i + split]
        ret_lst.append(dict(
            (string.cleanup(string.cleanhtml(x)), string.cleanup(string.cleanhtml(y))) for x, y in zip(title, line)))
    return ret_lst
