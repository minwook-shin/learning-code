import cssutils

sheet = cssutils.parseString('body { color: red }')
print("cssText1: ", sheet.cssText)

sheet = cssutils.parseString('body { color: red }')
cssutils.ser.prefs.indent = 4 * ' '
cssutils.ser.prefs.lineNumbers = True
print("cssText2: ", sheet.cssText)

from cssutils import css

sheet = css.CSSStyleSheet()
sheet.cssText = u'@import url(example.css) tv;'
print(sheet.cssText)

style = css.CSSStyleDeclaration()
style['color'] = 'red'
style_rule = css.CSSStyleRule(selectorText=u'body', style=style)
sheet.add(style_rule)
print(sheet.cssText)

print(sheet.cssText)
sheet.cssRules[0].media.appendMedium('print')
print(sheet.cssText)

sheet.cssRules[1].selectorList.appendSelector('a')
print(sheet.cssText)