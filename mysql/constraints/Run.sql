use harvest;

CALL addForeignKeyIfNotExists(
    'Run', 
    'FK_Run_Scraper_scraperId',
    'scraperId',
    'Scraper(scraperId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');

CALL addIndexIfNotExists('Run', 'IX_Run_startTimestamp', 'startTimestamp');
