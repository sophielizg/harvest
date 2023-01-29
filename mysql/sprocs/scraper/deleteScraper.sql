use harvest;

DROP PROCEDURE IF EXISTS deleteScraper;

DELIMITER $$

CREATE PROCEDURE deleteScraper(
    IN scraperIdIn INT,
    IN createTransaction BOOL
) BEGIN
    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        IF createTransaction THEN
            ROLLBACK;
        ELSE
            ROLLBACK TO deleteScraper;
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT deleteScraper;
    END IF;

    UPDATE Request req
    INNER JOIN Run r ON r.runId = req.runId
    SET 
        req.parentRequestId = NULL,
        req.originatorRequestId = NULL
    WHERE r.scraperId = scraperIdIn;

    DELETE FROM Scraper WHERE scraperId = scraperIdIn;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;