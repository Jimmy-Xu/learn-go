#!/bin/bash

# 查找/sys/xxx/ttyUSB*/dev


for sysdevpath in $(find /sys/bus/usb/devices/usb*/ -name dev | grep ttyUSB) 
do
    syspath="${sysdevpath%/dev}" # 去掉结尾的dev
    echo "syspath   : ${syspath}"
    eval "$(udevadm info -q property --export -p $syspath)" # 查询设备属性
    [[ -z "$ID_SERIAL" ]] && continue
    echo -e "ID_SERIAL: ${ID_SERIAL}"
    echo -e "${ID_VENDOR_ID}:${ID_MODEL_ID} ${DEVNAME}"
    echo -e "$(lsusb | grep ${ID_VENDOR_ID}:${ID_MODEL_ID})"
    echo
done
