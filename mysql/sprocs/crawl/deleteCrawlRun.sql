use harvest;

DROP PROCEDURE IF EXISTS deleteCrawlRun;

DELIMITER $$

CREATE PROCEDURE deleteCrawlRun(
    IN crawlRunIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE relatedCrawlId INT;
    DECLARE totalLeftInQueue INT;

    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO deleteCrawlRun;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT deleteCrawlRun;
    END IF;

    SELECT crawlId INTO relatedCrawlId FROM CrawlStatus
    WHERE crawlRunId = crawlRunIdIn;

    UPDATE Request r 
    INNER JOIN Scrape s ON r.scrapeId = s.scrapeId
    SET r.createdByRequestId = NULL
    WHERE s.crawlRunId = crawlRunIdIn;

    DELETE rs.* FROM Result rs
    INNER JOIN Request rq ON rs.requestId = rq.requestId
    INNER JOIN Scrape s ON rq.scrapeId = s.scrapeId
    WHERE s.crawlRunId = crawlRunIdIn;

    DELETE e.* FROM Error e
    INNER JOIN Request rq ON e.requestId = rq.requestId
    INNER JOIN Scrape s ON rq.scrapeId = s.scrapeId
    WHERE s.crawlRunId = crawlRunIdIn;

    SELECT SUM(cs.queued) INTO totalLeftInQueue FROM CrawlStatus cs
    LEFT JOIN Scrape s ON s.scrapeId = cs.scrapeId
    WHERE s.crawlRunId = crawlRunIdIn;

    CALL updateCrawlStatus(crawlIdIn, NULL, totalLeftInQueue, 0, 0, 0);

    DELETE cs.* FROM CrawlStatus cs
    LEFT JOIN Scrape s ON s.scrapeId = cs.scrapeId
    WHERE s.crawlRunId = crawlRunIdIn;

    DELETE rq.* FROM RequestQueue rq
    LEFT JOIN Scrape s ON s.scrapeId = rq.scrapeId
    WHERE s.crawlRunId = crawlRunIdIn;

    DELETE FROM CrawlRun
    WHERE crawlRunId = crawlRunIdIn;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;