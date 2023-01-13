use harvest;

DROP PROCEDURE IF EXISTS stopCrawl;

DELIMITER $$

CREATE PROCEDURE stopCrawl(
    IN crawlIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE currentCrawlRunId INT;

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
        SAVEPOINT stopCrawl;
    END IF;

    SELECT crawlRunId INTO currentCrawlRunId FROM CrawlRun
    WHERE crawlId = crawlIdIn 
    ORDER BY startTimestamp DESC
    LIMIT 1;

    UPDATE CrawlRun SET
        endTimestamp = NOW()
    WHERE crawlRunId = currentCrawlRunId;

    DELETE FROM RequestQueue
    WHERE crawlId = crawlIdIn AND isInitialRequest = 0;

    SET @numDequeued := ROW_COUNT();
    CALL updateCrawlStatus(crawlIdIn, NULL, -1 * @numDequeued, 0, 0, 0);

    UPDATE Crawl SET
        running = 0
    WHERE crawlId = crawlIdIn;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;