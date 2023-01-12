use harvest;

DROP PROCEDURE IF EXISTS addForeignKeyIfNotExists;

DELIMITER $$

CREATE PROCEDURE addForeignKeyIfNotExists(
    tableName VARCHAR(255),
    constraintName VARCHAR(255),
    columnName VARCHAR(255),
    referenceKey VARCHAR(255),
    onDeleteAndOnUpdate VARCHAR(255)
) BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
        WHERE TABLE_SCHEMA = DATABASE()
        AND CONSTRAINT_TYPE = 'FOREIGN KEY'
        AND CONSTRAINT_NAME = constraintName)
    THEN

        SET @sql := CONCAT('ALTER TABLE ', tableName,
                           ' ADD CONSTRAINT  ', constraintName,
                           ' FOREIGN KEY (', columnName, ')',
                           ' REFERENCES ', referenceKey,
                           ' ', onDeleteAndOnUpdate);
        PREPARE stmt FROM @sql;
        EXECUTE stmt;

    END IF;
END $$

DELIMITER ;