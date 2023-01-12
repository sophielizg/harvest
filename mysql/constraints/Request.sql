use harvest;

CALL addForeignKeyIfNotExists(
    'Request', 
    'FK_Request_Scrape_scrapeId',
    'scrapeId',
    'Scrape(scrapeId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Request', 
    'FK_Request_Request_requestId',
    'createdByRequestId',
    'Request(requestId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addIndexIfNotExists(
    'Request',
    'IX_Request_requestHash_visitedTimestamp', 
    'requestHash, visitedTimestamp');
