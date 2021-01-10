from django.db import models


class Segment(models.Model):
    width = models.IntegerField()
    height = models.IntegerField()
    square = models.FloatField(blank=True, null=True)
    created = models.DateTimeField(auto_now_add=True)
    deleted = models.DateTimeField(auto_now=True)
    defect = models.BooleanField(default=False)
    description = models.TextField(blank=True)
    active = models.BooleanField(default=True)
    color = models.ForeignKey(
        'segments.Color',
        on_delete=models.SET_NULL,
        related_name='segments',
        null=True,
    )
    order_number = models.ForeignKey(
        'segments.OrderNumber',
        on_delete=models.SET_NULL,
        related_name='segments',
        null=True,
    )
    rack = models.ForeignKey(
        'segments.Rack',
        on_delete=models.SET_NULL,
        related_name='segments',
        null=True,
    )

    class Meta:
        db_table = 'segment'
        verbose_name = 'Отрезок'
        verbose_name_plural = 'Отрезки'

    def __str__(self):
        return f'{self.color} - {self.width} - {self.height}'


class Rack(models.Model):
    name = models.CharField('Расположение', max_length=15)

    class Meta:
        verbose_name = 'Расположение'
        verbose_name_plural = 'Расположения'

    def __str__(self):
        return self.name


class OrderNumber(models.Model):
    name = models.CharField('Номер заказа', max_length=15)

    class Meta:
        verbose_name = 'Номер заказа'
        verbose_name_plural = 'Номера заказов'

    def __str__(self):
        return self.name


class ColorType(models.Model):
    name = models.CharField('Фактура', max_length=15)

    class Meta:
        verbose_name = 'Фактура'
        verbose_name_plural = 'Фактуры'

    def __str__(self):
        return self.name


class Color(models.Model):
    name = models.CharField('Цвет', max_length=30)
    type = models.ForeignKey(
        ColorType,
        on_delete=models.CASCADE,
        related_name='colors',
    )

    class Meta:
        verbose_name = 'Цвет'
        verbose_name_plural = 'Цвета'

    def __str__(self):
        return f'{self.type.name} - {self.name}'
