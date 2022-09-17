# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class TSPersonalStatItem(scrapy.Item):
    mission_id = scrapy.Field()
    found = scrapy.Field()
    http_status = scrapy.Field()
    source = scrapy.Field()
    nick = scrapy.Field()
    content = scrapy.Field()
    updated_at = scrapy.Field()

class GaijinPersonalItem(scrapy.Item):
    mission_id = scrapy.Field()
    found = scrapy.Field()
    http_status = scrapy.Field()
    source = scrapy.Field()
    nick = scrapy.Field()
    content = scrapy.Field()
    updated_at = scrapy.Field()