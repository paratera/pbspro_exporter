FROM centos:7 as builder

ENV GOROOT /usr/lib/golang
ENV GOPATH /go

WORKDIR /go

RUN  set -xeuo && \
    curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo && \
    curl -o /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo && \
    yum -y install golang && \
    mkdir -p $GOPATH/src/github.com/paratera && \
    mkdir -p $GOPATH/src/golang.org/x 

RUN set -xeuo && \
    yum groupinstall "Development Tools" -y && \
    yum install -y sudo tar wget openssh-server openssh-clients openssl openssl-devel && \
    yum install -y gcc make rpm-build libtool hwloc-devel libX11-devel libXt-devel libedit-devel libical-devel ncurses-devel perl postgresql-devel python-devel tcl-devel tk-devel swig expat-devel libXext libXft autoconf automake && \
    yum install -y expat libedit postgresql-server python sendmail tcl tk libical 

RUN set -xeuo && \
    git clone https://github.com/PBSPro/pbspro.git && \
    cd pbspro && \
    ./autogen.sh && \
    ./configure --prefix=/opt/pbspro && \
    make -j4 && \
    make install

RUN set -xeuo && \
    cd $GOPATH/src/golang.org/x && \
    git clone https://github.com/golang/tools.git && \
    git clone https://github.com/golang/lint.git && \
    git clone https://github.com/golang/sys.git && \
    git clone https://github.com/golang/crypto.git && \
    go get github.com/alecthomas/template && \
    go get github.com/alecthomas/template/parse && \
    go get github.com/alecthomas/units && \
    go get github.com/beorn7/perks/quantile && \
    go get github.com/golang/protobuf/proto && \
    go get github.com/juju/errors && \
    go get github.com/matttproud/golang_protobuf_extensions/pbutil && \
    go get github.com/siddontang/go/log && \
    go get github.com/paratera/go_pbspro/qstat && \
    go get github.com/paratera/go_pbspro/utils && \
    go get gopkg.in/alecthomas/kingpin.v2 && \
    go get github.com/prometheus/client_golang/prometheus && \
    go get github.com/prometheus/client_golang/prometheus/promhttp && \
    go get github.com/prometheus/client_model/go && \
    go get github.com/prometheus/common/expfmt && \
    go get github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg && \
    go get github.com/prometheus/common/log && \
    go get github.com/prometheus/common/model && \
    go get github.com/prometheus/common/version && \
    go get github.com/prometheus/procfs && \
    go get github.com/prometheus/procfs/internal/util && \
    go get github.com/prometheus/procfs/nfs && \
    go get github.com/prometheus/procfs/xfs && \
    go get github.com/sirupsen/logrus 


RUN set -xeuo && \
    mkdir -p $GOPATH/src/github.com/paratera && \
    cd $GOPATH/src/github.com/paratera && \
    git clone https://github.com/paratera/pbspro_exporter  && \
    cd pbspro_exporter && \
    go build


FROM centos:7 as running

COPY --from=builder /go/src/github.com/paratera/pbspro_exporter/pbspro_exporter /usr/bin/pbspro_exporter
COPY --from=builder /opt/pbspro/lib/libpbs.so.0 /usr/lib/libpbs.so.0

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
