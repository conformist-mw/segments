from django.db import models


class OrderNumber(models.Model):
    name = models.CharField('Номер заказа', max_length=15, unique=True)

    class Meta:
        verbose_name = 'Номер заказа'
        verbose_name_plural = 'Номера заказов'

    def __str__(self):
        return self.name
