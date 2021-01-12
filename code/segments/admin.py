from django.contrib import admin
from django.db.models import Count

from .models import Color, ColorType, OrderNumber, Rack, Segment


@admin.register(Segment)
class SegmentAdmin(admin.ModelAdmin):
    list_display = ['width', 'height', 'square', 'color', 'rack']
    list_select_related = ['rack', 'color', 'color__type']
    list_filter = ['rack', 'color__type']


@admin.register(Color)
class ColorAdmin(admin.ModelAdmin):
    list_display = ['name', 'type']
    list_filter = ['type']


@admin.register(ColorType)
class ColorTypeAdmin(admin.ModelAdmin):
    list_display = ['name', 'colors_count']

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
    list_display = ['name', 'segments_count']

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
