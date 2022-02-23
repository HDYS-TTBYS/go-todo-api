FROM postgres:latest
ARG DB_LANG=C
RUN localedef -i $DB_LANG -c -f UTF-8 -A /usr/share/locale/locale.alias $DB_LANG.UTF-8
ENV LANG $DB_LANG.utf8

# 初期化用 sql ファイルを、コンテナ内の所定のディレクトリにコピー
COPY ./init/*.sql /docker-entrypoint-initdb.d/
