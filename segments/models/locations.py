from django.db import models


class Rack(models.Model):
    name = models.CharField('Расположение', max_length=15)

    class Meta:
        verbose_name = 'Расположение'
        verbose_name_plural = 'Расположения'

    def __str__(self):
        return self.name
