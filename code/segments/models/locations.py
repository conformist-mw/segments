from django.db import models


class Company(models.Model):
    name = models.CharField('Компания', max_length=30, unique=True)
    slug = models.SlugField(max_length=50)
    image_url = models.CharField(max_length=255, blank=True)

    class Meta:
        verbose_name = 'Компания'
        verbose_name_plural = 'Компании'

    def __str__(self):
        return self.name


class Section(models.Model):
    name = models.CharField('Название цеха', max_length=30)
    slug = models.SlugField(max_length=50)
    company = models.ForeignKey(
        to=Company,
        verbose_name='Компания',
        related_name='sections',
        on_delete=models.CASCADE,
    )
    excluded_colors = models.ManyToManyField(
        to='segments.ColorType',
        verbose_name='Исключить фактуры',
        related_name='sections',
        blank=True,
    )

    class Meta:
        verbose_name = 'Цех'
        verbose_name_plural = 'Цеха'
        unique_together = ('name', 'company')

    def __str__(self):
        return f'{self.company.name} — {self.name}'


class Rack(models.Model):
    name = models.CharField('Расположение', max_length=30)
    section = models.ForeignKey(
        to=Section,
        verbose_name='Цех',
        related_name='racks',
        on_delete=models.CASCADE,
        null=True,  # delete it after migration
    )

    class Meta:
        verbose_name = 'Расположение'
        verbose_name_plural = 'Расположения'
        unique_together = ('name', 'section')

    def __str__(self):
        return self.name
