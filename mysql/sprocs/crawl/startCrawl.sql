use harvest;

DROP PROCEDURE IF EXISTS startCrawl;

DELIMITER $$

CREATE PROCEDURE startCrawl(
    IN crawlIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT startCrawl;
    END IF;

    INSERT INTO CrawlRun
        (crawlId, startTimestamp)
    VALUES
        (crawlIdIn, NOW());

    UPDATE Crawl SET
        running = 1
    WHERE crawlId = crawlIdIn;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;