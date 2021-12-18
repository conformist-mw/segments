from django.db.models import Q
from django_filters import rest_framework as filters
from segments.models import Segment


class SegmentFilter(filters.FilterSet):
    active = filters.BooleanFilter(field_name='active')
    color = filters.CharFilter(field_name='color__slug')
    type = filters.CharFilter(field_name='color__type__slug')
    width = filters.NumberFilter(field_name='width', method='size_filter')
    height = filters.NumberFilter(field_name='height', method='size_filter')
    order = filters.CharFilter(
        field_name='order_number__name',
        lookup_expr='icontains',
    )

    class Meta:
        model = Segment
        fields = ['active', 'color', 'type', 'width', 'height', 'order']

    def size_filter(self, queryset, name, value):
        return queryset.filter(Q(width__gte=value) | Q(height__gte=value))
