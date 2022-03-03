import datetime
import json
import random
import time
from multiprocessing import Process, Queue

import pika
from scrapy import crawler
from scrapy.utils.project import get_project_settings
from twisted.internet import reactor

from crawler.spiders.gaijin import GaijinSpider
from crawler.spiders.gaijin_cloudflare import GaijinCloudflareSpider
from crawler.spiders.thunderskill import ThunderSkillSpider


def f(q, spider, query_id, nick):
    try:
        runner = crawler.CrawlerRunner(get_project_settings())
        deferred = runner.crawl(spider, nick=nick, query_id=query_id)
        deferred.addBoth(lambda _: reactor.stop())
        reactor.run()
        q.put(None)
    except Exception as e:
        q.put(e)


def run_spider(spider, nick: str, query_id: str):
    q = Queue()
    p = Process(target=f, args=(q, spider, query_id, nick))
    p.start()
    result = q.get()
    p.join()

    if result is not None:
        raise result


def random_sleep_sec():
    return random.randrange(1, 10)


def callback(ch, method, properties, body):
    print("Received signal, start crawling")
    query_json = json.loads(body)
    print(query_json)
    begin = datetime.datetime.now()
    for i in query_json['target']:
        if i == 'thunder_skill':
            run_spider(spider=ThunderSkillSpider, nick=query_json['nickname'], query_id=query_json['query_id'])
        elif i == 'gaijin' and 'slow_mode' in query_json and query_json['slow_mode']:
            run_spider(spider=GaijinCloudflareSpider, nick=query_json['nickname'], query_id=query_json['query_id'])
        else:
            run_spider(spider=GaijinSpider, nick=query_json['nickname'], query_id=query_json['query_id'])
    end = datetime.datetime.now()
    sec = random_sleep_sec()
    print("Crawl finished, Spend %d seconds, sleep %d seconds. " % ((end - begin).seconds, sec))
    time.sleep(sec)


if __name__ == '__main__':
    connection = pika.BlockingConnection(pika.ConnectionParameters(get_project_settings().get("MQ_SOURCE")))
    channel = connection.channel()

    channel.queue_declare( queue='crawler')
    channel.basic_consume(queue='crawler',
                          auto_ack=True,
                          on_message_callback=callback)
    print("Start the query task...")
    channel.start_consuming()
