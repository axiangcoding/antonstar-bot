import json
import time
from datetime import datetime

import scrapy

from crawler.items import GaijinPersonalItem
from util import string


class GaijinSpider(scrapy.Spider):
    name = 'gaijin'
    allowed_domains = ['gaijin.net']

    def start_requests(self):
        nick = getattr(self, "nick", None)
        if nick is not None:
            url = f'https://warthunder.com/zh/community/userinfo/?nick={nick}'
            yield scrapy.Request(url, method='POST')

    # # mock
    # def start_requests(self):
    #     nick = getattr(self, "nick", None)
    #     if nick is not None:
    #         url = f'http://localhost:8888/api/v1/war_thunder/mock.html'
    #         yield scrapy.Request(url, method='GET')

    def parse(self, response):
        nick = getattr(self, "nick", None)
        query_id = getattr(self, "query_id", None)
        item = GaijinPersonalItem()
        item['query_id'] = query_id
        item['nick'] = nick
        item['source'] = 'gaijin'
        item['created_at'] = datetime.utcnow().strftime('%Y-%m-%d %H:%M:%S')
        item['http_status'] = response.status
        # TODO: 整理数据为json格式
        user_info_node = response.xpath("//div[@class='user-info']")
        nick = user_info_node.xpath("normalize-space(//li[@class='user-profile__data-nick']/text())").extract_first()
        banned_div = user_info_node.xpath("//div[@class='user-profile__data-nick--banned']").extract()
        banned = len(banned_div) != 0
        user_profile_list = user_info_node.xpath(
            "//li[@class='user-profile__data-item']/text() ").extract()
        register_date = user_info_node.xpath(
            "normalize-space(//li[@class='user-profile__data-regdate']/text())").extract_first()
        content = {
            'nick': nick,
            'banned': banned,
            'title': string.cleanup(user_profile_list[0]),
            'level': extract_level(string.cleanup(user_profile_list[1])),
            'register_date': extract_register_date(register_date)
        }
        item['content'] = json.dumps(content)
        # print(content)
        yield item


def extract_register_date(date_str: str):
    return time.strftime("%Y-%m-%d", time.strptime(date_str.split(" ")[1], "%d.%m.%Y"))


def extract_level(level_str: str):
    return level_str.replace("等级", "")
