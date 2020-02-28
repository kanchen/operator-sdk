FROM golang:1.14

RUN apt-get update 
RUN apt-get install curl 
RUN curl -fsSL https://get.docker.com -o get-docker.sh 
RUN sh get-docker.sh 
RUN apt-get install mercurial 
RUN apt-get install apt-transport-https 
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg |  apt-key add - 
RUN echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list 
RUN apt-get update 
RUN apt-get install kubectl 
RUN curl -LO https://github.com/operator-framework/operator-sdk/releases/download/v0.15.2/operator-sdk-v0.15.2-x86_64-linux-gnu
RUN chmod +x operator-sdk-v0.15.2-x86_64-linux-gnu && mkdir -p /usr/local/bin/ && cp operator-sdk-v0.15.2-x86_64-linux-gnu /usr/local/bin/operator-sdk && rm operator-sdk-v0.15.2-x86_64-linux-gnu

