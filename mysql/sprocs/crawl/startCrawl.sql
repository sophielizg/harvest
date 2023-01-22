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
        ELSE
            ROLLBACK TO startCrawl;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT startCrawl;
    END IF;

    INSERT INTO CrawlRun
        (crawlId, running, startTimestamp)
    VALUES
        (crawlIdIn, 1, NOW());

    UPDATE RequestQueue SET
        scrapeId = NULL
    WHERE isInitialRequest = 1;

    SET @numQueued := ROW_COUNT();
    CALL updateCrawlStatus(crawlIdIn, NULL, @numQueued, 0, 0, 0);

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;