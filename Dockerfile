# Dockerfile
FROM jenkins/jenkins:lts

USER root

# Install Go
RUN apt-get update && \
    apt-get install -y golang && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

USER jenkins
