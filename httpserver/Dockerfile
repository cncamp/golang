FROM ubuntu
ENV MY_SERVICE_PORT=80
ENV MY_SERVICE_PORT1=80
ENV MY_SERVICE_PORT2=80
ENV MY_SERVICE_PORT3=80
LABEL multi.label1="value1" multi.label2="value2" other="value3"
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver