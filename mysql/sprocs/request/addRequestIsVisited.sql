use harvest;

DROP PROCEDURE IF EXISTS addRequestIsVisited;

DELIMITER $$

CREATE PROCEDURE addRequestIsVisited(
    IN scrapeIdIn INT,
    IN requestHashIn VARCHAR(16),
    IN requestIn BLOB,
    IN createdByRequestIdIn INT
) BEGIN
    INSERT INTO Request
        (scrapeId, requestHash, visitedTimestamp, request, createdByRequestId)
    VALUES
        (scrapeIdIn, requestHashIn, NOW(), requestIn, createdByRequestIdIn);
    SELECT LAST_INSERT_ID() AS requestId;
END $$

DELIMITER ;