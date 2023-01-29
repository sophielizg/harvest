use harvest;

CALL addForeignKeyIfNotExists(
    'ParserTag', 
    'FK_ParserTag_Parser_parserId',
    'parserId',
    'Parser(parserId)',
    'ON DELETE CASCADE ON UPDATE CASCADE');

CALL addIndexIfNotExists('ParserTag', 'IX_ParserTag_name', 'name');
