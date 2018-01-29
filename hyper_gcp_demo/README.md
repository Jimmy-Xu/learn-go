
```
$ go build -o hyper main.go


$ export HYPER_ACCESS_KEY="XKV35VLYL9OXCExxxxxx3XXB"
$ export HYPER_SECRET_KEY="4yz1K1V93GXCNBp5oBBDqtxxxxxxm3WpNKBB"

$ ./hyper create
server: tcp://gcp-us-central1.hyper.sh:6443
1: create pod pod-1 : 1.09s
2: create pod pod-2 : 1.09s
3: create pod pod-3 : 1.06s
4: create pod pod-4 : 1.07s
...
100: create pod pod-100 : 1.07s


$ ./hyper list  
server: tcp://gcp-us-central1.hyper.sh:6443
1: pod-1 Running
2: pod-2 Running
3: pod-3 Running
4: pod-4 Running
...
100: pod-100 Running

```