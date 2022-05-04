FROM alpine
ADD micro-demo /micro-demo
ENTRYPOINT [ "/micro-demo" ]
