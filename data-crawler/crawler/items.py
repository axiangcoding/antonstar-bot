# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class TSPersonalStatItem(scrapy.Item):
    query_id = scrapy.Field()
    http_status = scrapy.Field()
    source = scrapy.Field()
    nick = scrapy.Field()
    content = scrapy.Field()
    created_at = scrapy.Field()

class GaijinPersonalItem(scrapy.Item):
    query_id = scrapy.Field()
    http_status = scrapy.Field()
    source = scrapy.Field()
    nick = scrapy.Field()
    content = scrapy.Field()
    created_at = scrapy.Field()