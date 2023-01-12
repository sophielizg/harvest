use harvest;

DROP PROCEDURE IF EXISTS addIndexIfNotExists;

DELIMITER $$

CREATE PROCEDURE addIndexIfNotExists(
    IN tableName VARCHAR(255),
    IN indexName VARCHAR(255),
    IN indexColumns VARCHAR(255)
) BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM INFORMATION_SCHEMA.STATISTICS
        WHERE TABLE_SCHEMA = DATABASE()
        AND TABLE_NAME = tableName
        AND INDEX_NAME = indexName)
    THEN

        SET @sql := CONCAT('ALTER TABLE ', tableName,
                           ' ADD INDEX  ', indexName,
                           ' (', indexColumns, ')');
        PREPARE stmt FROM @sql;
        EXECUTE stmt;

    END IF;
END $$

DELIMITER ;