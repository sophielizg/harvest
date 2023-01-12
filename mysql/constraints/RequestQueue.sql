use harvest;

CALL addForeignKeyIfNotExists(
    'RequestQueue', 
    'FK_RequestQueue_Scrape_scrapeId',
    'scrapeId',
    'Scrape(scrapeId)',
    'ON DELETE RESTRICT ON UPDATE CASCADE');
