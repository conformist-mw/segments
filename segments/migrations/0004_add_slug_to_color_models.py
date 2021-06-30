# Generated by Django 3.1.12 on 2021-06-28 19:45

from django.db import migrations, models
from slugify import slugify


def generate_slugs_for_names(apps, schema_editor):
    Color = apps.get_model('segments', 'Color')
    ColorType = apps.get_model('segments', 'ColorType')
    for color in Color.objects.all():
        color.slug = slugify(color.name, to_lower=True)
        color.save()
    for color_type in ColorType.objects.all():
        color_type.slug = slugify(color_type.name, to_lower=True)
        color_type.save()


class Migration(migrations.Migration):

    dependencies = [
        ('segments', '0003_add_company_and_section'),
    ]

    operations = [
        migrations.AddField(
            model_name='color',
            name='slug',
            field=models.SlugField(max_length=45, null=True),
        ),
        migrations.AddField(
            model_name='colortype',
            name='slug',
            field=models.SlugField(max_length=25, null=True),
        ),
        migrations.RunPython(
            generate_slugs_for_names, migrations.RunPython.noop,
        ),
    ]