
init-db:
	@echo "Initialize $(MYSQL_DATABASE)"
	@mysql -u root -e "DROP DATABASE IF EXISTS $(MYSQL_DATABASE);"
	@mysql -u root -e "CREATE DATABASE IF NOT EXISTS $(MYSQL_DATABASE) DEFAULT CHARSET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;"

run:
	go run main.go
