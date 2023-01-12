use harvest;

DROP PROCEDURE IF EXISTS updateCrawl;

DELIMITER $$

CREATE PROCEDURE updateCrawl(
    IN crawlIdIn INT
) BEGIN
    DECLARE `_rollback` BOOL DEFAULT 0;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET `_rollback` = 1;

    START TRANSACTION;

    UPDATE Request SET
        createdByRequestId = NULL
    WHERE crawlId = crawlIdIn;

    DELETE FROM Crawl WHERE rawlId = crawlIdIn;

    IF `_rollback` THEN
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;
END $$

DELIMITER ;