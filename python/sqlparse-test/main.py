import sqlparse


sql = 'select * from db; select * from db_slave;'
print(sqlparse.split(sql))


sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=False, keyword_case='lower'))
sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=False, keyword_case='upper'))

sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=True, keyword_case='lower'))
sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=True, keyword_case='upper'))

sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=True, identifier_case='lower'))
sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=True, identifier_case='upper'))

sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=True, identifier_case='lower',keyword_case='lower'))
sql = 'select * from db where id in (select id from db);'
print(sqlparse.format(sql, reindent=True, identifier_case='upper',keyword_case='upper'))


sql = 'delete from db where id = 0'
parsed = sqlparse.parse(sql)
print(parsed)
for i in parsed:
    print(i)