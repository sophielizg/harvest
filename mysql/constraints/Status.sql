use harvest;

CALL addForeignKeyIfNotExists(
    'Status', 
    'FK_Status_Run_runId',
    'runId',
    'Run(runId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');

CALL addForeignKeyIfNotExists(
    'Status', 
    'FK_Status_RunnerQueue_runnerId',
    'runnerId',
    'RunnerQueue(runnerId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');

CALL addIndexIfNotExists(
    'Status',
    'IX_Status_lastUpdatedTimestamp',
    'lastUpdatedTimestamp');
