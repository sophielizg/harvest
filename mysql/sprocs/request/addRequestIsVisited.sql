use harvest;

DROP PROCEDURE IF EXISTS addRequestIsVisited;

DELIMITER $$

CREATE PROCEDURE addRequestIsVisited(
    IN requestIdIn UNSIGNED INT,
    IN crawlRunIdIn INT,
    IN requestIn BLOB,
    IN createdByRequestIdIn INT
) BEGIN
    INSERT INTO Request
        (requestId, crawlRunId, visitedTimestamp, request, createdByRequestId)
    VALUES
        (requestIdIn, crawlRunIdIn, NOW(), requestIn, createdByRequestIdIn)
    ON DUPLICATE KEY UPDATE
        request = requestIn,
        createdByRequestId = createdByRequestIdIn;
END $$

DELIMITER ;