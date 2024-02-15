FROM ubuntu:22.04
RUN apt update
RUN apt install -y curl
RUN apt install -y libfuse2
RUN curl https://storage.googleapis.com/packages.viam.com/apps/viam-server/viam-server-stable-aarch64.AppImage -o viam-server && chmod 755 viam-server
CMD ./viam-server --appimage-extract-and-run -config /config/viam.json
