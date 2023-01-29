use harvest;

CALL addForeignKeyIfNotExists(
    'Request', 
    'FK_Request_Run_runId',
    'runId',
    'Run(runId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Request', 
    'FK_Request_Request_parentRequestId',
    'parentRequestId',
    'Request(requestId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Request', 
    'FK_Request_Request_originatorRequestId',
    'originatorRequestId',
    'Request(requestId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addIndexIfNotExists(
    'Request',
    'IX_Request_visitedTimestamp', 
    'visitedTimestamp');
