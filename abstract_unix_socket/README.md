 # show criu.sock
$ ss -xlp | grep criu.sock
u_str  LISTEN     0      128    @criu.sock 249839

$ # connect to the abstract socket
$ socat - ABSTRACT-CONNECT:criu.sock
hello
hello
world
world
