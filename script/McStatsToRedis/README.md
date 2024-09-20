## McStatsToRedis
* 此脚本用于将Minecraft服务器的玩家统计数据存储到redis的有序集合里。
* This script is used to store player statistic data in Minecraft server to sorted set of Redis.
* 此脚本主要面向Linux，在Windows里使用可能有问题。
* This script is for Linux, problems may occurred when using in Windows.
* 先修改config.json里的配置，如果Redis有密码，可以在config.json的redis对象里加一个password字段，值为密码字符串。
* Edit configurations in config.json first, if your Redis has password, you can add a "password" kay-value pair in "redis" object in config.json, the value is your password string.
* 先进入python虚拟环境，再执行脚本即可，输入下面的命令。
* Enter python environment first, then run the script with the command below.
* `source myenv/bin/activate`
* `python3 McStatsToRedis.py config.json`
* 也可以不进入虚拟环境执行代码。
* You can also run the command without entering the environment.
* `myenv/bin/python3 McStatsToRedis.py config.json`
* config.json里的`sumed_stats`和`custom_stats`代表统计的信息类型，`mc_servers`代表统计的服务器，`mc_servers.path`代表服务器的根路径，`mc_servers.db`代表将数据存储到的数据库编号。
* The `sumed_stats` and `custom_stats` in config.json stand for the data type to be counted. `mc_servers` is servers to be counted. `mc_servers.path` and `mc_servers.db` are the root path of the server and the db in where data is store.
