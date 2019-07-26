from grab.spider import Spider, Task


class Main(Spider):
    def task_generator(self):
        for lang in 'python', 'java', 'go':
            url = 'https://www.google.com/search?q=%s' % lang
            yield Task('result', url=url, lang=lang)

    def task_result(self, grab, task):
        print('{}: {}'.format(task.lang, grab.doc('//*[@id="rhs_block"]').text()))


bot = Main(thread_number=3)
bot.run()