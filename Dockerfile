FROM golang:1.12.1

WORKDIR /work
RUN curl -O https://bootstrap.pypa.io/get-pip.py \
  && python get-pip.py \
  && pip install awscli \
  && pip install --upgrade awscli \
  && curl -sL https://deb.nodesource.com/setup_11.x | bash - \
  && apt-get install -y python2.7-dev nodejs apt-transport-https ca-certificates gnupg-agent software-properties-common \
  && curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - \
  && add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/debian \
    $(lsb_release -cs) \
    stable" \
  && apt-get update \
  && apt-get install -y docker-ce-cli \
  && npm install -g serverless \
  && pip install aws-sam-cli \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && echo "export DOCKER_HOST='tcp://host.docker.internal:2375'" >> ~/.bashrc

EXPOSE 3000

WORKDIR /go

ENTRYPOINT ["/bin/bash"]