#使うDockerイメージ
FROM mysql:8.0
ENV TZ=Asia/Tokyo
#ポートを開ける
EXPOSE 3306
#MySQL設定ファイルをイメージ内にコピー
ADD ./my.cnf /etc/mysql/conf.d/my.cnf
#docker runに実行される
CMD ["mysqld"]
