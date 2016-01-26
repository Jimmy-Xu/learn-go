#!/bin/bash
#############################################
# build docker from source
# 1.get docker source by "go get"
# 2.append pprof to docker/docker/docker.go
# 3.run docker with -cpuprofile
#############################################



CUR_DIR="$(cd `dirname $0`; pwd)"

DOCKER_BASE="${GOPATH}/src/github.com/docker/docker"
DOCKER_SRC="${DOCKER_BASE}/docker/docker.go"
DOCKERFILE="${DOCKER_BASE}/Dockerfile"
DOCKER_BUNDLE="${DOCKER_BASE}/bundles/latest/binary/docker"
PPROF_FILE="/tmp/docker_cpu.pprof"
PDF_FILE="${HOME}/docker_callgraph.pdf"

echo "DOCKER_BASE: ${DOCKER_BASE}"
echo "DOCKER_SRC:  ${DOCKER_SRC}"



function quit(){
  echo "$1"
  exit 1
}

function show_usage(){
  cat <<EOF

[usage]
  ./docker.sh <action>
  <action>: build|run|graph

-----------------------------------------------------------------
1.build docker source

  ./docker.sh build <option>  //build docker source with pprof

<option>
  --tag : list all tags
  [tag] : build specified tag

example:
  ./docker.sh build v1.9.1
  ./docker.sh build v1.10.0-rc1
  ./docker.sh build master

-----------------------------------------------------------------
2.run docker daemon

  ./docker.sh run         //run docker daemon with pprof

-----------------------------------------------------------------
3.generate callgraph pdf

  ./docker.sh graph       //convert ${PPROF_FILE} -> ./${PDF_FILE}

EOF
  exit 1
}


function do_run(){
  echo "-------------------------------------"
  echo " > run docker daemon"
  sudo service docker stop
  touch ${PPROF_FILE}
  sudo ${DOCKER_BUNDLE} daemon -D --cpuprofile=${PPROF_FILE}
}

function do_build(){

  if [ "$1" == "--tag" ];then
    cd ${DOCKER_BASE} && git tag --list && cd -
    exit 0
  fi

  TGT_BRANCH=$1
  echo "-------------------------------------"
  echo " > get docker source code"
  go get github.com/docker/docker
  if [ "${TGT_BRANCH}" != "" ];then
    cd ${DOCKER_BASE}
    ORIG_BRANCH=$(git branch  | grep "*" | awk '{print $2}')
    echo " > CURRENT BRANCH: ${ORIG_BRANCH} -> TARGET BRANCH: ${TGT_BRANCH} "
    if [ "${TGT_BRANCH}" != "${ORIG_BRANCH}" ] ;then
      git checkout master -f
      if [[ "${ORIG_BRANCH}" != "master" ]];then
        echo " > delete old branch '${ORIG_BRANCH}'"
        git branch -D ${ORIG_BRANCH}
      fi
      if [ "${TGT_BRANCH}" != "master" ];then
        echo " > start checkout branch to ${TGT_BRANCH}"
        git reset --hard HEAD && git checkout ${TGT_BRANCH} -f && git checkout -b ${TGT_BRANCH} && cd -
        if [ $? -ne 0 ];then
          quit "checkout to branch ${TGT_BRANCH} failed"
        fi
      fi
    else
      echo "current branch is target branch, skip switch branch"
    fi
  fi

  echo " > check docker.go"
  if [ ! -f ${DOCKER_SRC} ];then
      quit "${DOCKER_SRC} not found"
  fi


  echo "-------------------------------------"
  echo " > append pprof to docker/docker/docker.go"

  grep import ${DOCKER_SRC} -A1 | grep "runtime/pprof" >/dev/null 2>&1
  if [ $? -ne 0 ];then
      sed -i '/import/a \\t"runtime/pprof"' ${DOCKER_SRC}
  fi

  grep 'func main() {' ${DOCKER_SRC} -A1 | grep 'var cpuprofile =' >/dev/null 2>&1
  if [ $? -ne 0 ];then
      sed -i "/func main() {/a \\\tvar cpuprofile = \"${PPROF_FILE}\"\n\tcommonFlags.FlagSet.StringVar(&cpuprofile,[]string{\"-cpuprofile\"}, \"${PPROF_FILE}\", \"write cpu profile to file\")" ${DOCKER_SRC}
  fi

  grep 'flag.Parse' ${DOCKER_SRC} -A1 | grep 'if cpuprofile != "" {' >/dev/null 2>&1
  if [ $? -ne 0 ];then
      sed -i '/flag.Parse/a \\tif cpuprofile != "" {\n\t\tf, err := os.Create(cpuprofile)\n\t\tif err != nil { fmt.Println("Error: ", err) }\n\t\tpprof.StartCPUProfile(f)\n\t\tdefer pprof.StopCPUProfile()\n\t}' ${DOCKER_SRC}
  fi

  echo "-------------------------------------"
  echo " > modify Dockerfile"
  echo " > create ${DOCKER_BASE}/sources.list "
  cat <<EOF > ${DOCKER_BASE}/sources.list
      deb http://mirrors.163.com/ubuntu/ trusty main restricted universe multiverse
      deb http://mirrors.163.com/ubuntu/ trusty-security main restricted universe multiverse
      deb http://mirrors.163.com/ubuntu/ trusty-updates main restricted universe multiverse
      deb http://mirrors.163.com/ubuntu/ trusty-proposed main restricted universe multiverse
      deb http://mirrors.163.com/ubuntu/ trusty-backports main restricted universe multiverse
      deb-src http://mirrors.163.com/ubuntu/ trusty main restricted universe multiverse
      deb-src http://mirrors.163.com/ubuntu/ trusty-security main restricted universe multiverse
      deb-src http://mirrors.163.com/ubuntu/ trusty-updates main restricted universe multiverse
      deb-src http://mirrors.163.com/ubuntu/ trusty-proposed main restricted universe multiverse
      deb-src http://mirrors.163.com/ubuntu/ trusty-backports main restricted universe multiverse
      deb http://security.ubuntu.com/ubuntu trusty-security main
EOF
  echo " > add sources.list to ${DOCKERFILE}"
  grep 'ADD sources.list /etc/apt/sources.list' ${DOCKERFILE} > /dev/null 2>&1
  if [ $? -ne 0 ];then
      sed -i "/FROM ubuntu:/a ADD sources.list /etc/apt/sources.list" ${DOCKERFILE}
  fi
  echo " > add proxy to ${DOCKERFILE}"
  grep 'ENV http_proxy' ${DOCKERFILE} > /dev/null 2>&1
  if [ $? -ne 0 ];then
      DOCKER_HOST=$(ip route | grep docker0 | awk '{print $NF}')
      sed -i "/FROM ubuntu:/a ENV http_proxy http://${DOCKER_HOST}:8118\nENV https_proxy https://${DOCKER_HOST}:8118\nENV no_proxy localhost,127.0.0.0/8,::1,mirrors.163.com" ${DOCKERFILE}
  fi
  echo " > update go-md2man 1.0.3->1.0.4"
  sed -i 's#v1.0.3 https://github.com/cpuguy83/go-md2man.git#v1.0.4 https://github.com/cpuguy83/go-md2man.git#' ${DOCKERFILE}
  echo " > update blackfriday 1.2->1.4"
  sed -i 's#v1.2 https://github.com/russross/blackfriday.git#v1.4 https://github.com/russross/blackfriday.git#' ${DOCKERFILE}

  echo " > after modified"
  cd ${DOCKER_BASE} && git diff --exit-code && cd -
  sleep 1

  echo "-------------------------------------"
  echo " > check docker daemon"
  sudo service docker status| grep running > /dev/null 2>&1
  if [ $? -ne 0 ];then
      echo " > start docker daemon..."
      sudo service docker start
      if [ $? -ne 0 ];then
          quit "docker daemon not running..."
      fi
  fi

  echo "-------------------------------------"
  echo " > start build docker from source"
  cd ${DOCKER_BASE} && make && cd -

  echo "-------------------------------------"
  echo " > show build result"
  ls -l --color ${DOCKER_BUNDLE}
}

function do_graph(){
  echo "-------------------------------------"
  echo " > generate call graph to pdf format"
  go tool pprof --pdf ${DOCKER_BUNDLE} ${PPROF_FILE} > ${PDF_FILE}
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
    do_build "$2"
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
