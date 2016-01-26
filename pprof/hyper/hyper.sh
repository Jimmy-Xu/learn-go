#!/bin/bash
#############################################
# build hyperd from source
# 1.get hyperd source
# 2.append pprof to hyperhq/hyper/hyperd.go
# 3.run hyperd with -cpuprofile
#############################################

CUR_DIR="$(cd `dirname $0`; pwd)"

HYPER_BASE="${GOPATH}/src/github.com/hyperhq/hyper"
HYPER_SRC="${HYPER_BASE}/hyperd.go"
HYPERD_BIN="${HYPER_BASE}/hyperd"
PPROF_FILE="/tmp/hyperd_cpu.pprof"
PDF_FILE="${HOME}/hyperd_callgraph.pdf"

echo "HYPER_BASE: ${HYPER_BASE}"
echo "HYPER_SRC:  ${HYPER_SRC}"

function quit(){
  echo "$1"
  exit 1
}

function show_usage(){
  cat <<EOF
usage:
  ./hyper.sh <action>
  <action>: build|run|graph
eg:
  ./hyper.sh build    //build hyperd with pprof
  ./hyper.sh run      //run hyperd daemon with pprof
  ./hyper.sh graph    //convert ${PPROF_FILE} -> ./${PDF_FILE}
EOF
  exit 1
}


function do_run(){
  echo "-------------------------------------"
  echo " > run docker daemon"
  sudo service hyperd stop
  touch ${PPROF_FILE}
  sudo ${HYPERD_BIN} --nondaemon -cpuprofile=${PPROF_FILE}
}

function do_build(){
  echo "-------------------------------------"
  echo " > get hyper source code"
  #go get github.com/hyperhq/hyper
  if [ ! -d ${HYPER_BASE} ];then
    echo " > start clone repo"
    git clone https://github.com/hyperhq/hyper.git ${HYPER_BASE}
  else
    echo " > start pull newest code"
    cd ${HYPER_BASE} && git pull && cd -
  fi

  echo " > check hyperd.go"
  if [ ! -f ${HYPER_SRC} ];then
      quit "${HYPER_SRC} not found"
  fi

  echo "-------------------------------------"
  echo " > append pprof to ${HYPER_SRC}"
  #1
  grep import ${HYPER_SRC} -A1 | grep "runtime/pprof" >/dev/null 2>&1
  if [ $? -ne 0 ];then
      sed -i '/import/a \\t"runtime/pprof"' ${HYPER_SRC}
  fi
  #2
  grep 'func main() {' ${HYPER_SRC} -A1 | grep 'var cpuprofile =' >/dev/null 2>&1
  if [ $? -ne 0 ];then
      sed -i '/func main() {/a \\tvar cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")' ${HYPER_SRC}
  fi
  #3
  grep 'flag.Parse' ${HYPER_SRC} -A1 | grep 'if \*cpuprofile != "" {' >/dev/null 2>&1
  if [ $? -ne 0 ];then
      sed -i '/flag.Parse/a \\tif *cpuprofile != "" {\n\t\tf, err := os.Create(*cpuprofile)\n\t\tif err != nil { fmt.Println("Error: ", err) }\n\t\tpprof.StartCPUProfile(f)\n\t\tdefer pprof.StopCPUProfile()\n\t}' ${HYPER_SRC}
  fi
  echo " > check after modified"
  cd ${HYPER_BASE} && git diff --exit-code && cd -


  echo "-------------------------------------"
  echo " > start build hyperd"
  cd ${HYPER_BASE} && \
  ./autogen.sh && \
  ./configure --without-xen && \
  make && cd -
  if [ $? -ne 0 ];then
    quit "build hyperd failed,quit"
  fi

  echo "-------------------------------------"
  echo " > show build result"
  ls -l --color ${HYPERD_BIN}
}

function do_graph(){
  echo "-------------------------------------"
  echo " > generate call graph to pdf format"
  CMD_LINE="go tool pprof --pdf ${HYPERD_BIN} ${PPROF_FILE} > ${PDF_FILE}"
  $CMD_LINE
  if [ ! -s ${PDF_FILE} ];then
    quit " > failed"
  else
    echo "Done"
    ls -l ${PDF_FILE}
  fi
}

##### main #####

if [ $# -eq 0 ];then
  show_usage
fi

case "$1" in
  build)
    do_build
    ;;
  run)
    do_run
    ;;
  graph)
    do_graph
    ;;
  *)
    show_usage
    ;;
esac
echo "Done"
