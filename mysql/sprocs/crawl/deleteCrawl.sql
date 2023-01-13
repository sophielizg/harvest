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
        END IF;
        RESIGNAL;
    END;

    IF createTransaction THEN
        START TRANSACTION;
    ELSE
        SAVEPOINT deleteCrawl;
    END IF;

    UPDATE Request SET
        createdByRequestId = NULL
    WHERE crawlId = crawlIdIn;

    DELETE FROM Crawl WHERE crawlId = crawlIdIn;

    IF createTransaction THEN
        COMMIT;
    END IF;
END $$

DELIMITER ;