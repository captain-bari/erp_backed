#!/bin/bash 

while true
do
  /app/erp >> /app/log/out.log 2>>/app/log/err.log
  sleep 5
done