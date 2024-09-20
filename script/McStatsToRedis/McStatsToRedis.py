import sys
import redis
import json
import os

def store(data, records_name, method, player_name, server_name):
    for record_name in records_name:
        full_record_name = 'minecraft:'+record_name
        if full_record_name in data:
            value = method(data[full_record_name])
            r.zadd(server_name+':'+record_name, {player_name: value})

if len(sys.argv) != 2:
    print("用法：python 脚本名.py 配置文件路径.json")
    sys.exit(1)

with open(sys.argv[1], 'r', encoding='utf-8') as file:
    config = json.load(file)

r = redis.Redis(**config['redis'])
r.flushdb()

for server in config['mc_servers']:
    server_path = server['path']
    server_name = server['name']
    usercache_path = os.path.join(server_path, 'usercache.json')
    with open(usercache_path, 'r', encoding='utf-8') as file:
        usercache = json.load(file)
    hashtable_name = server_name+':uuid_name'
    for obj in usercache:
        r.hset(server_name+':name_uuid', obj['name'], obj['uuid'])
        r.hset(hashtable_name, obj['uuid'], obj['name'])
    stats_folder = os.path.join(server_path, 'world/stats/')
    for filename in os.listdir(stats_folder):
        if filename.endswith('.json'):
            uuid = os.path.splitext(filename)[0]
            if r.hexists(hashtable_name, uuid):
                player_name = r.hget(hashtable_name, uuid)
                stats_file = os.path.join(stats_folder, filename)
                with open(stats_file, 'r', encoding='utf-8') as file:
                    player_data = json.load(file)['stats']
                    store(
                        player_data, 
                        config['sumed_stats'], 
                        lambda x: sum(x.values()), 
                        player_name,
                        server_name,
                    )
                    if 'minecraft:custom' in player_data:
                        store(
                            player_data['minecraft:custom'], 
                            config['custom_stats'], 
                            lambda x: x, 
                            player_name,
                            server_name,
                        )
