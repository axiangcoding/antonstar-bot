import json
import time

from src.core import api


def collect_api_info(needs: list):
    batch = dict()
    batch['timestamp'] = time.time()
    if 'mission' in needs:
        batch['mission'] = api.mission_json()
    if 'map_obj' in needs:
        batch['map_obj'] = api.map_obj_json()
    if 'map_info' in needs:
        batch['map_info'] = api.map_info_json()
    if 'gamechat' in needs:
        batch['gamechat'] = api.gamechat()
    if 'indicators' in needs:
        batch['indicators'] = api.indicators()
    if 'state' in needs:
        batch['state'] = api.state()
    return json.dumps(batch, ensure_ascii=False)
