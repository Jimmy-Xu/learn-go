```
$ go build

$ sudo ./thinkgos --action=get
modbusRTUMaster => 2020/07/17 23:12:30 [D]: sending [01 03 00 01 00 01 d5 ca]
modbusRTUMaster => 2020/07/17 23:12:30 [D]: received [01 03 02 00 04 b9 87]
INFO[0000] []uint16{0x4}  

$ sudo ./thinkgos --action=set --port=4
modbusRTUMaster => 2020/07/17 23:12:01 [D]: sending [01 06 00 01 00 04 d9 c9]
modbusRTUMaster => 2020/07/17 23:12:01 [D]: received [01 06 00 01 00 04 d9 c9]
INFO[0000] write single register OK  
```