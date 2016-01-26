#!/bin/bash
# This shell will do the following:
# 1 donwload and untar source code of golang to ~/go
# 2 set GO env(GOROOT=~/go, GOPATH=~/gopath)
# 3 modify ${VCS_SRC}, improve go get,  add --progress to git command line
# 4 build golang build image by Dockerfile
# 5 start new container to build golang (about 10min)
#
# result is $GOROOT/bin/go

GO_VER="1.5.3"
DOCKER_IMAGE="xjimmyshcn/golang:${GO_VER}"
BASE_DIR=$(cd `dirname $0`; pwd)

function quit(){
    echo $1
    exit 1
}

function download_go(){
    cd ${HOME}
    echo "-----------------------------------------------"
    echo " > get go${GO_VER}.linux-amd64.tar.gz "
    echo "-----------------------------------------------"
    wget -c https://storage.googleapis.com/golang/go${GO_VER}.linux-amd64.tar.gz
    tar xzvf go${GO_VER}.linux-amd64.tar.gz

    echo "-----------------------------------------------"
    echo " > get go1.5.3.src.tar.gz "
    echo "-----------------------------------------------"
    wget -c https://storage.googleapis.com/golang/go${GO_VER}.src.tar.gz
    tar xzvf go${GO_VER}.src.tar.gz
}

function update_vcs(){
    VCS_SRC="${GOROOT}/src/cmd/go/vcs.go"
    echo "-----------------------------------------------"
    echo " > modify ${VCS_SRC} "
    if [ -f ${VCS_SRC} ];then

        echo "-----------------------------------------------"
        echo " > update git command line parameter "
        sed -i 's/createCmd:   \[\]string{"clone {repo}/createCmd:   \[\]string{"clone --progress {repo}/' ${VCS_SRC}

        grep 'cmd.Stderr = &buf' ${VCS_SRC} -A1 | grep 'cmd.Stdout = os.Stdout' >/dev/null 2>&1
        if [ $? -ne 0 ];then
            echo "-----------------------------------------------"
            echo " > update stdout and stderr "
            sed -i '/cmd.Stderr = &buf/a \\tcmd.Stdout = os.Stdout\n\tcmd.Stderr = os.Stderr' ${VCS_SRC}
        fi
        echo "-----------------------------------------------"
        echo " > after modified "
        grep -E '(= os.Std|--progress)' ${VCS_SRC}
        echo "-----------------------------------------------"
    else
       quit "${VCS_SRC} not found,quit"
    fi
}

function build_image(){
  cd ${BASE_DIR}
  echo "-----------------------------------------------"
  echo " > start build image ${DOCKER_IMAGE}"
  docker build -t ${DOCKER_IMAGE} .

  if [ $? -ne 0 ];then
      quit "docker build failed!"
  fi
}

##### main #####
_SHELL=$(echo $SHELL | awk -F"/" '{print $NF}')
case "${_SHELL}" in
    zsh)        SH_RC="${HOME}/.zshrc" ;;
    sh|bash)    SH_RC="${HOME}/.bashrc" ;;
    *) quit "shell '${_SHELL}' not supported";;
esac

#ensure ~/go, ~/gopath
mkdir -p ${HOME}/gopath

#ensure GOROOT and GOPATH
grep "^export GOROOT=${HOME}/go" ${SH_RC} >/dev/null 2>&1 || echo "export GOROOT=$HOME/go" >> ${SH_RC}
grep "^export GOPATH=${HOME}/gopath" ${SH_RC} >/dev/null 2>&1 || echo "export GOPATH=$HOME/gopath" >> ${SH_RC}
grep '^export PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}' ${SH_RC} >/dev/null 2>&1 || echo 'export PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}' >> ${SH_RC}

#check ENV
echo "-----------------------------------------------"
grep -i "^export .*go" ${SH_RC}

#download golang source
download_go

#update ${GOROOT}/src/cmd/go/vcs.go
update_vcs

#build xjimmyshcn/golang:${GO_VER}
build_image

#start golang container
#docker run -it --rm -v $GOROOT:/go -v $GOPATH:/root/gopath -e GOPATH=/root/gopath -e GOROOT_BOOTSTRAP=/usr/local/go ${DOCKER_IMAGE}
docker run -it --rm -v $GOROOT:/go -v $GOPATH:/root/gopath ${DOCKER_IMAGE}

echo "-----------------------------------------------"
echo " >build result:"
ls -l --color $GOROOT/bin/go
echo "Done"
