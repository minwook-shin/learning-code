import xlsxwriter

workbook = xlsxwriter.Workbook('test2.xlsx')
worksheet = workbook.add_worksheet()

data = [10, 20, 30, 40, 50]
worksheet.write_column('A1', data)

chart = workbook.add_chart({'type': 'line'})
chart.add_series({'values': '=Sheet1!$A$1:$A$5'})

worksheet.insert_chart('C1', chart)

workbook.close()
