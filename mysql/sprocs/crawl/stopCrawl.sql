use harvest;

DROP PROCEDURE IF EXISTS stopCrawl;

DELIMITER $$

CREATE PROCEDURE stopCrawl(
    IN crawlIdIn INT
) BEGIN
    DECLARE `_rollback` BOOL DEFAULT 0;
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET `_rollback` = 1;

    START TRANSACTION;

    SELECT @crawlRunId := crawlRunId FROM CrawlRun
    WHERE crawlId = crawlIdIn 
    ORDER BY startTimestamp DESC
    LIMIT 1 FOR UPDATE;

    UPDATE CrawlRun SET
        endTimestamp = NOW()
    WHERE crawlRunId = @crawlRunId;

    DELETE FROM RequestQueue
    WHERE crawlId = crawlIdIn AND isInitialRequest = 0;

    SELECT @numDequeued := ROW_COUNT();
    CALL updateCrawlStatus(crawlIdIn, NULL, -1 * @numDequeued, 0, 0, 0);

    UPDATE Crawl SET
        running = 0
    WHERE crawlId = crawlIdIn;

    IF `_rollback` THEN
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;
END $$

DELIMITER ;