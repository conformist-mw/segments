from django.db import models


class ColorTypeManager(models.Manager):
    def filter_by(self, section):
        return super().get_queryset().exclude(
            pk__in=section.excluded_colors.values_list('id', flat=True),
        )


class ColorType(models.Model):
    name = models.CharField('Фактура', max_length=15, unique=True)
    slug = models.SlugField(max_length=25, unique=True)

    objects = ColorTypeManager()

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
    slug = models.SlugField(max_length=45)

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
