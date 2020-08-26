FROM alpine

# RUN, COPY, ADD , FROM

WORKDIR /app
# ADD vs COPY
COPY . .

#* Comandos de Linux (alpine)
RUN ls /app && chmod +x myapp.sh && apk add --update bash curl wget

CMD [ "./myapp.sh" ]
ENTRYPOINT [ "bash" ]


