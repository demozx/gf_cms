FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
ADD resource                $WORKDIR/
ADD config.yaml             $WORKDIR/
ADD ./temp/linux_amd64/gfCms $WORKDIR/gfCms
RUN chmod +x $WORKDIR/gfCms

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./gfCms
