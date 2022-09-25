from django.contrib import admin
from django.db.models import Count

from .models import (
    Color,
    ColorType,
    Company,
    OrderNumber,
    Rack,
    Section,
    Segment,
    SegmentOld,
)


@admin.register(SegmentOld)
class SegmentOldAdmin(admin.ModelAdmin):
    list_display = [
        'width', 'height', 'square', 'color', 'rack',
        'order_number', 'active',
    ]


@admin.register(Segment)
class SegmentAdmin(admin.ModelAdmin):
    list_display = [
        'id', 'width', 'height', 'square', 'color', 'rack',
        'order_number', 'active', 'created',
    ]
    list_select_related = ['rack', 'color', 'color__type', 'order_number']
    list_filter = ['active', 'rack', 'color__type', 'rack__section']
    readonly_fields = ['order_number']


@admin.register(Company)
class CompanyAdmin(admin.ModelAdmin):
    list_display = ['name']
    prepopulated_fields = {'slug': ['name']}


@admin.register(Section)
class SectionAdmin(admin.ModelAdmin):
    list_display = ['name']
    prepopulated_fields = {'slug': ['name']}


@admin.register(Color)
class ColorAdmin(admin.ModelAdmin):
    list_display = ['name', 'type']
    list_filter = ['type']
    prepopulated_fields = {'slug': ['name']}


@admin.register(ColorType)
class ColorTypeAdmin(admin.ModelAdmin):
    list_display = ['name', 'colors_count']
    prepopulated_fields = {'slug': ['name']}

    def get_queryset(self, request):
        return (
            super().get_queryset(request)
            .annotate(colors_count=Count('colors'))
        )

    def colors_count(self, obj):
        return obj.colors_count

    colors_count.short_description = 'Количество цветов'


@admin.register(Rack)
class RackAdmin(admin.ModelAdmin):
    list_display = ['name', 'segments_count', 'section']
    list_filter = ['section', 'section__company']

    def get_queryset(self, request):
        return (
            super().get_queryset(request)
            .annotate(segments_count=Count('segments'))
            .order_by('name')
        )

    def segments_count(self, obj):
        return obj.segments_count

    segments_count.short_description = 'Количество отрезков'


@admin.register(OrderNumber)
class OrderNumberAdmin(admin.ModelAdmin):
    pass
