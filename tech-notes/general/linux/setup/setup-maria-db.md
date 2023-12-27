---
tags:
  - linux-setup
---
1. `sudo apt install mariadb-server`
2. `sudo mysql_secure_installation`. This will take you through a series of prompts where you can make some changes to your MariaDB installation’s security options. The first prompt will ask you to enter the current database root password. Since you have not set one up yet, press ENTER to indicate “none”.
3. The next prompt asks you whether you’d like to set up a database root password. On Ubuntu, the root account for MariaDB is tied closely to automated system maintenance, so we should not change the configured authentication methods for that account. Doing so would make it possible for a package update to break the database system by removing access to the administrative account. Type N and then press ENTER.
4. From there, you can press Y and then ENTER to accept the defaults for all the subsequent questions. This will remove some anonymous users and the test database, disable remote root logins, and load these new rules so that MariaDB immediately implements the changes you have made.
5. `sudo mariadb`
6. `GRANT ALL ON *.* TO 'admin'@'localhost' IDENTIFIED BY 'admin' WITH GRANT OPTION;`
7. `FLUSH PRIVILEGES;`
8. `mysqladmin -u admin -p`
