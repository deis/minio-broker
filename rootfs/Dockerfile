FROM quay.io/deis/base:v0.3.5

# Add user and group
RUN adduser --system \
	--shell /bin/bash \
	--disabled-password \
	--home /opt/minio-broker \
	--group \
	minio-broker

COPY . /

# Fix some permission since we'll be running as a non-root user
RUN chown -R minio-broker:minio-broker /opt/minio-broker \
    && chmod +x /opt/minio-broker/sbin/*

USER minio-broker
WORKDIR /opt/minio-broker/sbin/
CMD ["./minio-broker"]

ARG VERSION
ARG BUILD_DATE
ENV VERSION $VERSION
ENV BUILD_DATE $BUILD_DATE
ENV PATH $PATH:/opt/discapi/sbin