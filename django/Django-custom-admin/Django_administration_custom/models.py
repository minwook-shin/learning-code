from django.db import models


# Create your models here.
class DB(models.Model):
    test = models.CharField(max_length=15)
    permission = models.CharField(max_length=15,choices=(('test', "test"),('real', "real")),blank=True, null=True)

    def __str__(self):
        return self.test
