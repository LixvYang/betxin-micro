#生成的表名
MODEL=$2

# 数据库配置
host=127.0.0.1
port=3306
dbname=betxin
username=root
passwd=123456

echo "开始创建库：$dbname 的表：$2"

sql2pb -field_style sql_pb -go_package ./pb -host ${host} -package pb -password ${passwd} -port ${port} -schema ${dbname} -service_name ${MODEL}srv -user $root -table=${MODEL} > ${MODEL}srv.proto