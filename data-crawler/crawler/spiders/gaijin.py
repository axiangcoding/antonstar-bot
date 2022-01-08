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
            url = f'https://warthunder.com/en/community/userinfo/?nick={nick}'
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
        user_info_node = response.xpath("//div[@class='user-info']")
        nick = user_info_node.xpath("normalize-space(//li[@class='user-profile__data-nick']/text())").extract_first()
        clan = user_info_node.xpath("//li[@class='user-profile__data-clan']/a/text()").extract_first()
        clan_url = user_info_node.xpath("//li[@class='user-profile__data-clan']/a/@href").extract_first()
        banned_div = user_info_node.xpath("//div[@class='user-profile__data-nick--banned']").extract()
        banned = len(banned_div) != 0
        user_profile_list = user_info_node.xpath(
            "//li[@class='user-profile__data-item']/text() ").extract()
        register_date = user_info_node.xpath(
            "normalize-space(//li[@class='user-profile__data-regdate']/text())").extract_first()

        user_stat = {}
        user_stat_list = \
            user_info_node.xpath(
                "//div[@class='user-stat__list-row user-stat__list-row--with-head']/ul[contains(@class, 'user-stat__list')]/li").extract()
        title_size = int(len(user_stat_list) / 4)
        temp_lst = zip_user_stat(user_stat_list, title_size)
        user_stat['ab'] = temp_lst[0]
        user_stat['rb'] = temp_lst[1]
        user_stat['sb'] = temp_lst[2]

        user_rate = {
            'aviation': {},
            'ground_vehicles': {},
            'fleet': {}
        }
        user_rate_temp_list = \
            user_info_node.xpath(
                "//div[@class='user-profile__stat user-stat user-stat--tabs']/div[1]/ul/li").extract()
        temp_lst = zip_user_stat(user_rate_temp_list, int(len(user_rate_temp_list) / 4))
        user_rate['aviation']['ab'] = temp_lst[0]
        user_rate['aviation']['rb'] = temp_lst[1]
        user_rate['aviation']['sb'] = temp_lst[2]

        user_rate_temp_list = \
            user_info_node.xpath(
                "//div[@class='user-profile__stat user-stat user-stat--tabs']/div[2]/ul/li").extract()
        temp_lst = zip_user_stat(user_rate_temp_list, int(len(user_rate_temp_list) / 4))
        user_rate['ground_vehicles']['ab'] = temp_lst[0]
        user_rate['ground_vehicles']['rb'] = temp_lst[1]
        user_rate['ground_vehicles']['sb'] = temp_lst[2]

        user_rate_temp_list = \
            user_info_node.xpath(
                "//div[@class='user-profile__stat user-stat user-stat--tabs']/div[3]/ul/li").extract()
        temp_lst = zip_user_stat(user_rate_temp_list, int(len(user_rate_temp_list) / 4))
        user_rate['fleet']['ab'] = temp_lst[0]
        user_rate['fleet']['rb'] = temp_lst[1]
        user_rate['fleet']['sb'] = temp_lst[2]

        content = {
            'nick': nick,
            'clan': clan,
            'clan_url': clan_url,
            'banned': banned,
            'title': string.cleanup(user_profile_list[0]),
            'level': extract_level(string.cleanup(user_profile_list[1])),
            'register_date': extract_register_date(register_date),
            'user_stat': user_stat,
            'user_rate': user_rate
        }
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
