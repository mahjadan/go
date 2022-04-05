FROM python:3.7-slim-buster

ADD . /src
WORKDIR /src

# Fix issue with JRE installation
RUN apt-get update && \
    apt-get install -y sudo && \
    apt-get install -y curl && \
    apt-get install -y openjdk-11-jre-headless && \
    apt-get clean;

RUN curl -sL firebase.tools | bash
RUN firebase --version

CMD [ "firebase", "--project=go-proto-poc", "emulators:start" ]
