use harvest;

DROP PROCEDURE IF EXISTS addResult;

DELIMITER $$

CREATE PROCEDURE addResult(
    IN crawlIdIn INT,
    IN scrapeIdIn INT,
    IN requestIdIn INT,
    IN parserIdIn INT,
    IN valueIn TEXT,
    IN createTransaction BOOL
) BEGIN
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO addResult;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT addResult;
    END IF;

    INSERT INTO Result
        (requestId, parserId, scrapedTimestamp, value)
    VALUES
        (requestIdIn, parserIdIn, NOW(), valueIn);

    CALL updateCrawlStatus(crawlIdIn, scrapeIdIn, 0, 1, 0, 0);
    
    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;