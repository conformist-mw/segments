from django.db.models import Count, Sum
from django.db.models.functions import Round
from rest_framework.pagination import LimitOffsetPagination
from rest_framework.viewsets import ModelViewSet

from segments.models import Company, Section, Segment

from .serializers import (
    CompanySerializer,
    SectionSerializer,
    SegmentSerializer,
)


class SegmentsPagination(LimitOffsetPagination):
    max_limit = 20


class CompanyViewSet(ModelViewSet):
    serializer_class = CompanySerializer
    pagination_class = None
    queryset = Company.objects.all()
    lookup_field = 'slug'


class SectionViewSet(ModelViewSet):
    serializer_class = SectionSerializer
    pagination_class = None
    lookup_field = 'slug'
    queryset = Section.objects.select_related('company').all()

    def get_queryset(self):
        return (
            self.queryset
            .filter(company__slug=self.kwargs['company_slug'])
            .annotate(segments_count=Count('racks__segments'))
            .annotate(
                square_sum=Round(
                    Sum('racks__segments__square', distinct=True),
                ),
            )
            .annotate(racks_count=Count('racks', distinct=True))
        )


class SegmentsViewSet(ModelViewSet):
    serializer_class = SegmentSerializer
    pagination_class = SegmentsPagination
    queryset = Segment.objects.select_related('color__type', 'rack').all()

    def get_queryset(self):
        return self.queryset.filter(
            rack__section__slug=self.kwargs['section_slug'],
            rack__section__company__slug=self.kwargs['company_slug'],
        )
