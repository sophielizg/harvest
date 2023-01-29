use harvest;

CALL addForeignKeyIfNotExists(
    'RunnerQueue', 
    'FK_RunnerQueue_Run_runId',
    'runId',
    'Run(runId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');
