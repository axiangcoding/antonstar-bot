# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import pymysql
from itemadapter import ItemAdapter
from scrapy.utils.project import get_project_settings


class MysqlPipeline:
    def __init__(self):
        settings = get_project_settings()
        self.conn = pymysql.connect(host=settings['MYSQL_HOST'], port=settings['MYSQL_PORT'],
                                    user=settings['MYSQL_USER'], password=settings['MYSQL_PWD'],
                                    database=settings['MYSQL_DB'], charset='utf8mb4',
                                    )
        self.cursor = self.conn.cursor()

    def process_item(self, item, spider):
        # sql语句
        insert_sql = """
                insert into crawler_data(query_id,http_status,source,nick,content,created_at) VALUES(%s,%s,%s,%s,%s,%s)
                """
        # 执行插入数据到数据库操作
        self.cursor.execute(insert_sql,
                            (item['query_id'], item['http_status'], item['source'], item['nick'], item['content'],
                             item['created_at']))
        # 提交，不进行提交无法保存到数据库
        self.conn.commit()

    def close_spider(self, spider):
        self.cursor.close()
        self.conn.close()
