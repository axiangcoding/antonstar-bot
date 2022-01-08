# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import pymysql
from itemadapter import ItemAdapter


class MysqlPipeline:
    def __init__(self):
        self.conn = pymysql.connect(host='127.0.0.1', port=3306,
                                    user='antonstar', password='AntonStarP@ssword.', database='anton_star',
                                    charset='utf8mb4',
                                    )
        self.cursor = self.conn.cursor()

    def process_item(self, item, spider):
        # sql语句
        insert_sql = """
                insert into crawler_data(http_status,source,nick,content,created_at) VALUES(%s,%s,%s,%s,%s)
                """
        # 执行插入数据到数据库操作
        self.cursor.execute(insert_sql, (item['http_status'], item['source'], item['nick'], item['content'],
                                         item['created_at']))
        # 提交，不进行提交无法保存到数据库
        self.conn.commit()

    def close_spider(self, spider):
        self.cursor.close()
        self.conn.close()
