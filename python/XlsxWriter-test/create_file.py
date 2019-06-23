import xlsxwriter

workbook = xlsxwriter.Workbook('test1.xlsx')
worksheet = workbook.add_worksheet()

test_tuple = (
    ['name', 'count'],
    ['test_name_1', 10],
    ['test_name_2', 20],
    ['test_name_3', 30],
)

row = 0
for name, count in test_tuple:
    worksheet.write(row, 0, name)
    worksheet.write(row, 1, count)
    row += 1

worksheet.write(row, 0, 'Total')
worksheet.write(row, 1, '=SUM(B2:B4)')

workbook.close()