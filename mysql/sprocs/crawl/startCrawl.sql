use harvest;

DROP PROCEDURE IF EXISTS startCrawl;

DELIMITER $$

CREATE PROCEDURE startCrawl(
    IN crawlIdIn INT
) BEGIN
    DECLARE `_rollback` BOOL DEFAULT 0;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET `_rollback` = 1;

    START TRANSACTION;

    INSERT INTO CrawlRun
        (crawlId, startTimestamp)
    VALUES
        (crawlIdIn, NOW());

    UPDATE Crawl SET
        running = 1
    WHERE crawlId = crawlIdIn;

    IF `_rollback` THEN
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;
END $$

DELIMITER ;