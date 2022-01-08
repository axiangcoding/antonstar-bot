import logging

from scrapy.crawler import CrawlerProcess
from scrapy.utils.project import get_project_settings

from crawler.spiders.gaijin import GaijinSpider
from crawler.spiders.thunderskill import ThunderSkillSpider
import pika


def run_spider(spider, nick: str):
    """会启用pipeline"""
    process = CrawlerProcess(get_project_settings())
    try:
        process.crawl(spider, nick=nick)
        process.start()
    except Exception as e:
        process.stop()
        logging.error("errorMsg:%s" % e.message)


def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)
    # name = 'gaijin'
    # nickname = 'OnTheRocks'
    # if name == 'thunderskill':
    #     run_spider(spider=ThunderSkillSpider, nick=nickname)
    # elif name == 'gaijin':
    #     run_spider(spider=GaijinSpider, nick=nickname)


if __name__ == '__main__':
    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = connection.channel()

    channel.queue_declare(queue='crawler')
    channel.basic_consume(queue='crawler',
                          auto_ack=True,
                          on_message_callback=callback)
    channel.start_consuming()
