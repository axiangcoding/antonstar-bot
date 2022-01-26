import json
import time
from datetime import datetime

import scrapy
import undetected_chromedriver as uc
from lxml import etree

from crawler.items import GaijinPersonalItem
from util.extract_gaijin import extract_user_info


class GaijinCloudflareSpider(scrapy.Spider):
    name = 'gaijin-cloudflare'
    allowed_domains = ['gaijin.net']
    start_urls = ['http://localhost']

    def __init__(self, **kwargs):
        super().__init__(**kwargs)
        options = uc.ChromeOptions()
        options.add_argument('--no-first-run --no-service-autorun --password-store=basic')
        self.driver = uc.Chrome(options=options, suppress_welcome=True, version_main=97)

    def parse(self, response):
        nick = getattr(self, "nick", None)
        query_id = getattr(self, "query_id", None)
        if nick is not None:
            url = f'https://warthunder.com/zh/community/userinfo/?nick={nick}'
            self.driver.get(url)
            time.sleep(10)
            text = self.driver.page_source
            self.driver.close()
            html = etree.HTML(text)
            item = GaijinPersonalItem()
            item['query_id'] = query_id
            item['nick'] = nick
            item['source'] = 'gaijin'
            item['updated_at'] = datetime.utcnow().strftime('%Y-%m-%d %H:%M:%S')
            item['http_status'] = 0

            user_not_found = html.xpath("//div[@class='user__unavailable-title']")
            if len(user_not_found) > 0:
                item['content'] = 'not found'
                item['found'] = False
            else:
                item['found'] = True
                content = extract_user_info(text)
                item['content'] = json.dumps(content, ensure_ascii=False)
            yield item
