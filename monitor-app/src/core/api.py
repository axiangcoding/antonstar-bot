import json
from json.decoder import JSONDecodeError
from urllib import parse

import requests

base_url = 'http://127.0.0.1:8111'

session = requests.Session()


def state():
    path = 'state'
    url = parse.urljoin(base_url, path)
    res = session.get(url)
    return res.json()


def mission_json():
    path = 'mission.json'
    url = parse.urljoin(base_url, path)
    res = session.get(url)
    return res.json()


def map_obj_json():
    path = 'map_obj.json'
    url = parse.urljoin(base_url, path)
    res = session.get(url)
    try:
        return json.loads(res.text)
    except JSONDecodeError:
        return None


def map_info_json():
    path = 'map_info.json'
    url = parse.urljoin(base_url, path)
    res = session.get(url)
    return res.json()


def gamechat(last_id=0):
    path = 'gamechat'
    url = parse.urljoin(base_url, path)
    res = session.get(url=url, params={
        'lastId': last_id
    })
    return json.loads(res.text)


def hudmsg(last_evt=0, last_dmg=0):
    path = 'hudmsg'
    url = parse.urljoin(base_url, path)
    res = session.get(url=url, params={
        'lastEvt': last_evt,
        'lastDmg': last_dmg
    })
    return res.json()


def indicators():
    path = 'indicators'
    url = parse.urljoin(base_url, path)
    res = session.get(url)
    return res.json()
