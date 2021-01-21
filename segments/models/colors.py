from django.db import models


class ColorType(models.Model):
    name = models.CharField('Фактура', max_length=15, unique=True)

    class Meta:
        verbose_name = 'Фактура'
        verbose_name_plural = 'Фактуры'

    def __str__(self):
        return self.name


class ColorManager(models.Manager):
    def get_queryset(self):
        return super().get_queryset().select_related('type')


class Color(models.Model):
    name = models.CharField('Цвет', max_length=30)
    type = models.ForeignKey(
        ColorType,
        on_delete=models.CASCADE,
        related_name='colors',
    )

    objects = ColorManager()

    class Meta:
        verbose_name = 'Цвет'
        verbose_name_plural = 'Цвета'
        unique_together = ('name', 'type')

    def __str__(self):
        return f'{self.type.name} - {self.name}'
