# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import json

import pymysql
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
        # update_crawler_data_sql = """
        # update crawler_data
        # set found=%s,http_status=%s,content=%s,updated_at=%s,status='done'
        # where query_id=%s and source=%s
        # """
        # # 执行插入数据到数据库操作
        # self.cursor.execute(update_crawler_data_sql,
        #                     (item['found'], item['http_status'], item['content'],
        #                      item['updated_at'], item['query_id'], item['source'],))
        # FIXME: mission的result里面，只会保留最后一个执行的内容，需要修改
        update_mission_sql = """
        update mission set result=%s, status=%s where mission_id=%s
        """
        self.cursor.execute(update_mission_sql, (item['content'], 'success', item['query_id']))
        # 提交，不进行提交无法保存到数据库
        self.conn.commit()
        # 如果找到用户并且是gaijin的数据源，则进行静态记录的更新
        if item['found']:
            if item['source'] == 'gaijin':
                content = json.loads(item['content'])
                find_wt_user_sql = """
                select id from game_user where nick=%s
                """
                self.cursor.execute(find_wt_user_sql, content['nick'])
                res = self.cursor.fetchall()

                # 如果找到记录，说明不是第一次获取到该玩家记录
                if len(res) != 0:
                    update_wt_user_sql = """
                    update game_user
                    set clan=%s, clan_url=%s, register_date=%s, level=%s, title=%s, banned=%s, updated_at=%s 
                    where nick=%s 
                    """
                    self.cursor.execute(update_wt_user_sql,
                                        (content['clan'], content['clan_url'],
                                         content['register_date'], content['level'], content['title'],
                                         content['banned'], item['updated_at'], content['nick']))
                else:
                    insert_wt_user_sql = """
                    insert into game_user(nick,clan,clan_url,register_date,level,title,banned,created_at,updated_at) 
                    values (%s,%s,%s,%s,%s,%s,%s,%s,%s)
                    """
                    self.cursor.execute(insert_wt_user_sql,
                                        (content['nick'], content['clan'], content['clan_url'],
                                         content['register_date'], content['level'], content['title'],
                                         content['banned'], item['updated_at'], item['updated_at']))
                self.conn.commit()
            if item['source'] == 'thunder_skill':
                content = json.loads(item['content'])
                update_wt_user_sql = """
                   update game_user set ts_ab_rate=%s, ts_rb_rate=%s, ts_sb_rate=%s 
                   where nick=%s 
                """
                self.cursor.execute(update_wt_user_sql,
                                    (content['a']['kpd'], content['r']['kpd'], content['s']['kpd'], content['nick']))
                self.conn.commit()

    def close_spider(self, spider):
        self.cursor.close()
        self.conn.close()
