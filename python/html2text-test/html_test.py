import html2text

html_2_text = html2text.HTML2Text()
html = """
<!DOCTYPE html>
<html>
<body>

<h1>First Heading</h1>
<h2>Second Heading</h2>
<h3>Third Heading</h3>
<p>Paragraph.</p>
<b>Bold.</b>
<a href="https://www.google.com">google</a>

<img src="google.jpg" alt="google.com">

<ul>
  <li>java</li>
  <li>python</li>
  <li>go</li>
</ul>

<ol>
  <li>java</li>
  <li>python</li>
  <li>go</li>
</ol>

<table>
  <tr>
    <th>id</th>
    <th>name</th> 
    <th>e-mail</th>
  </tr>
  <tr>
    <td>test_id1</td>
    <td>test1</td> 
    <td>test@test.com</td>
  </tr>
  <tr>
    <td>test_id2</td>
    <td>test2</td> 
    <td>another@test.com</td>
  </tr>
</table>

</body>
</html>"""

result = html_2_text.handle(html)
print(result)