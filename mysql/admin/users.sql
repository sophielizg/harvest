CREATE USER '$MYSQL_APP_USER'@'%' IDENTIFIED BY '$MYSQL_APP_PASSWORD';
CREATE USER '$MYSQL_BUILDER_USER'@'%' IDENTIFIED BY '$MYSQL_BUILDER_PASSWORD';

GRANT 
    ALTER, ALTER ROUTINE, CREATE, CREATE ROUTINE, DROP, INDEX, REFERENCES,
    EXECUTE, INSERT, SELECT, UPDATE, DELETE
ON harvest.* TO '$MYSQL_BUILDER_USER'@'%';

GRANT 
    EXECUTE, INSERT, SELECT, UPDATE, DELETE 
ON harvest.* TO '$MYSQL_APP_USER'@'%';
