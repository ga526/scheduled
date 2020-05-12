#!/usr/bin/env bash
nowday=$(date "+%Y%m%d%H%M");
rm -rf docker/step_*.txt;
echo '1'>docker/step_1.txt && go build main.go;
echo '2'>docker/step_2.txt && chmod -R 777 main;
echo '3'>docker/step_3.txt && ./main >>docker/run_log_${nowday}.txt 2>&1 ;
echo '4'>docker/step_success.txt;