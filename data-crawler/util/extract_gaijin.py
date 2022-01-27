import time

from lxml import etree

from util import string


def extract_user_info(page_html: str):
    html = etree.HTML(page_html)
    user_info_node = html.xpath("//div[@class='user-info']")[0]
    nick = user_info_node.xpath(
        "normalize-space(//li[@class='user-profile__data-nick']/text())")
    clan_list = user_info_node.xpath("//li[@class='user-profile__data-clan']/a/text()")
    clan_url_list = user_info_node.xpath("//li[@class='user-profile__data-clan']/a/@href")
    banned_div = user_info_node.xpath("//div[@class='user-profile__data-nick--banned']")
    register_date = user_info_node.xpath(
        "normalize-space(//li[@class='user-profile__data-regdate']/text())")

    user_profile_list = user_info_node.xpath(
        "//li[@class='user-profile__data-item']/text()")
    user_stat = {}
    user_stat_list = user_info_node.xpath(
        "//div[@class='user-stat__list-row user-stat__list-row--with-head']/ul[contains(@class, 'user-stat__list')]/li")
    text_user_stat_list = []
    for i in user_stat_list:
        text_user_stat_list.append(i.text)
    title_size = int(len(text_user_stat_list) / 4)
    temp_lst = zip_user_stat(text_user_stat_list, title_size)
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
            "//div[@class='user-profile__stat user-stat user-stat--tabs']/div[1]/ul/li")
    text_user_rate_list = []
    for i in user_rate_temp_list:
        text_user_rate_list.append(i.text)
    temp_lst = zip_user_stat(text_user_rate_list, int(len(text_user_rate_list) / 4))
    user_rate['aviation']['ab'] = temp_lst[0]
    user_rate['aviation']['rb'] = temp_lst[1]
    user_rate['aviation']['sb'] = temp_lst[2]

    user_rate_temp_list = \
        user_info_node.xpath(
            "//div[@class='user-profile__stat user-stat user-stat--tabs']/div[2]/ul/li")
    text_user_rate_list = []
    for i in user_rate_temp_list:
        text_user_rate_list.append(i.text)
    temp_lst = zip_user_stat(text_user_rate_list, int(len(text_user_rate_list) / 4))
    user_rate['ground_vehicles']['ab'] = temp_lst[0]
    user_rate['ground_vehicles']['rb'] = temp_lst[1]
    user_rate['ground_vehicles']['sb'] = temp_lst[2]

    user_rate_temp_list = \
        user_info_node.xpath(
            "//div[@class='user-profile__stat user-stat user-stat--tabs']/div[3]/ul/li")
    text_user_rate_list = []
    for i in user_rate_temp_list:
        text_user_rate_list.append(i.text)
    temp_lst = zip_user_stat(text_user_rate_list, int(len(text_user_rate_list) / 4))
    user_rate['fleet']['ab'] = temp_lst[0]
    user_rate['fleet']['rb'] = temp_lst[1]
    user_rate['fleet']['sb'] = temp_lst[2]

    res = {
        'nick': nick,
        'clan': '' if len(clan_list) == 0 else clan_list[0],
        'clan_url': '' if len(clan_url_list) == 0 else clan_url_list[0],
        'banned': len(banned_div) != 0,
        'register_date': extract_register_date(register_date),
        'title': string.cleanup(user_profile_list[0]),
        'level': extract_level(string.cleanup(user_profile_list[1])),
        'user_stat': user_stat,
        'user_rate': user_rate

    }
    return res


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
            (string.cleanup(x), string.cleanup(y)) for x, y in zip(title, line)))
    return ret_lst
