# * app (src) + Dockerfile => docker image build => docker image (exe)

# * docker image (exe) => docker container run => docker container (instance)

FROM ubuntu:20.04

COPY myapp.sh /home/

RUN chmod +x /home/myapp.sh && ls -l /home && \
    apt-get update && \
    apt-get install -y curl

CMD [ "/home/myapp.sh" ]
#ENTRYPOINT [ "executable" ]
