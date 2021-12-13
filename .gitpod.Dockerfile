# v1 https://community.gitpod.io/t/how-do-i-try-mongodb-with-gitpod/152
# v2 https://www.gitpod.io/blog/gitpodify/#mongodb

FROM gitpod/workspace-mongodb

RUN go get -u github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo

# RUN mkdir -p /tmp/mongodb && \
#     cd /tmp/mongodb && \
#     wget -qOmongodb.tgz https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu2004-5.0.2.tgz && \
#     tar xf mongodb.tgz && \
#     cd mongodb-* && \
#     sudo cp bin/* /usr/local/bin/ && \
#     rm -rf /tmp/mongodb && \
#     sudo mkdir -p /data/db && \
#     sudo chown gitpod:gitpod -R /data/db