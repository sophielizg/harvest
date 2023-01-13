use harvest;

DROP PROCEDURE IF EXISTS addError;

DELIMITER $$

CREATE PROCEDURE addError(
    IN crawlIdIn INT,
    IN scrapeIdIn INT,
    IN requestIdIn INT,
    IN parserIdIn INT,
    IN statusCodeIn INT,
    IN isMissngParseResultIn BOOL,
    IN responseIn TEXT,
    IN createTransaction BOOL
) BEGIN
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO addError;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT addError;
    END IF;

    INSERT INTO Error
        (requestId, parserId, scrapedTimestamp,
         statusCode, isMissngParseResult, response)
    VALUES
        (requestIdIn, parserIdIn, NOW(),
         statusCodeIn, isMissngParseResultIn, responseIn);

    CALL updateCrawlStatus(
        crawlIdIn, scrapeIdIn, 0, 0, 
        IF(isMissngParseResultIn, 0, 1), 
        IF(isMissngParseResultIn, 1, 0));
    
    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;