FROM alpine


VOLUME [ "/volumes" ]
# RUN, COPY, ADD , FROM

WORKDIR /app
# ADD vs COPY
COPY . .

COPY . /volumes/

#* Comandos de Linux (alpine)
RUN ls /app && chmod +x myapp.sh && apk add --update bash curl wget

CMD [ "./myapp.sh" ]
ENTRYPOINT [ "bash" ]

