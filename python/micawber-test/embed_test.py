import micawber
providers = micawber.bootstrap_basic()
embed_output = providers.parse_text('바로 보기 : \n https://www.youtube.com/watch?v=R6IT_f0XPT8')
print(embed_output)