# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import json
import logging

import requests
from scrapy.utils.project import get_project_settings


class MysqlPipeline:
    def process_item(self, item, spider):
        pattern = "{}/v1/crawler/callback?app_token={}"
        url = get_project_settings().get("API_SYSTEM_URL")
        token = get_project_settings().get("API_SYSTEM_SECRET")
        content = json.dumps(item["content"], ensure_ascii=False)
        resp = requests.post(url=pattern.format(url, token),
                             json={
                                 "crawler_data": content,
                                 'mission_id': item['mission_id'],
                                 'source': item['source'],
                             })
        if resp.status_code != 200:
            logging.error(resp)
