FROM docker.4pd.io/base-image-common:1.0
ARG PACKAGE_NAME=app-manager
WORKDIR /home/work

COPY ${PACKAGE_NAME} ./bin/${PACKAGE_NAME}
COPY entrypoint.sh .

RUN chmod 0755 entrypoint.sh
ENTRYPOINT ["/bin/sh", "./entrypoint.sh"]
