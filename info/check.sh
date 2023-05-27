#!/usr/bin/bash
function datetime() {
    date +"%Y-%m-%d %H:%M:%S"
}

while true; do
  flag=`ps -ef |grep mongodump |grep -Ev "color=auto"`
  flagRestore=`ps -ef | egrep -vw "grep|ansible|tail|tailf|less|vim|bash" | grep 'mongorestore' | wc -l`
  if [[ -z "$flag" ]]; then
      # 为空没有进程执行
      curl --header "Content-Type:application/json"  --data '{"level":"Info","msg":"mongodump执行结束"}'  -X POST http://123.60.223.82:60123/info
      break
  fi

  if [ $flagRestore -ne 3 ]; then
    curl --header "Content-Type:application/json"  --data '{"level":"Info","msg":"mongorestore执行结束一个"}'  -X POST http://123.60.223.82:60123/info
    break
  fi
  sleep 2
  echo $(datetime) "接着进程检查"
done