FROM ubuntu:22.04
RUN apt update
RUN apt install -y curl
RUN apt install -y libfuse2
COPY setup.sh ./setup.sh
CMD ./setup.sh
