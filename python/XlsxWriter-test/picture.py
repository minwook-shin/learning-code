import xlsxwriter

workbook = xlsxwriter.Workbook('test3.xlsx')
worksheet = workbook.add_worksheet()

worksheet.insert_image('A1', 'test.png')

workbook.close()