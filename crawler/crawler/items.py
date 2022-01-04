# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class TSPersonalStatItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    http_status = scrapy.Field()
    http_response = scrapy.Field()
    createAt = scrapy.Field()

class GaijinPersonalItem(scrapy.Item):
    http_status = scrapy.Field()
    http_response = scrapy.Field()
    createAt = scrapy.Field()