#!/usr/bin/env bash
echo '1'>docker/step_1.txt && go build main.go;
echo '2'>docker/step_2.txt && chmod -R 777 main;
echo '3'>docker/step_3.txt && ./main &;