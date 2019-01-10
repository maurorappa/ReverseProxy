FROM busybox
RUN mkdir -p /srv/log /srv/config
COPY ReverseProxy /srv/ReverseProxy
EXPOSE     80
ENTRYPOINT [ "/srv/ReverseProxy" ]
#CMD [ "-conf", "/srv/config/rp.conf" ]

