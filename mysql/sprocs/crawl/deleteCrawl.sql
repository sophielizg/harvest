use harvest;

DROP PROCEDURE IF EXISTS deleteCrawl;

DELIMITER $$

CREATE PROCEDURE deleteCrawl(
    IN crawlIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO deleteCrawl;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT deleteCrawl;
    END IF;

    UPDATE Request r
    INNER JOIN Scrape s ON r.scrapeId = s.scrapeId
    INNER JOIN CrawlRun cr ON s.crawlRunId = cr.crawlRunId
    SET r.createdByRequestId = NULL
    WHERE cr.crawlId = crawlIdIn;

    DELETE FROM Crawl WHERE crawlId = crawlIdIn;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;