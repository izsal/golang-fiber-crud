- make the database with docker like this and adjust
  docker run --rm -d --name atlas-demo -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e
  MYSQL_DATABASE=example mysql

- how to run migrate to database mysql
  to check different from folder entities "atlas migrate diff --env gorm"

- to run migrate to your database mysql "atlas migrate apply --env gorm -u "mysql://root:password@localhost:3306/example""
