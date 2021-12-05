from rest_framework.viewsets import ModelViewSet

from segments.models import Segment, Company, Section

from .serializers import SegmentSerializer, CompanySerializer, SectionSerializer


class CompanyViewSet(ModelViewSet):
    serializer_class = CompanySerializer
    queryset = Company.objects.all()
    lookup_field = 'slug'


class SectionViewSet(ModelViewSet):
    serializer_class = SectionSerializer
    lookup_field = 'slug'
    queryset = Section.objects.select_related('company').all()

    def get_queryset(self):
        return self.queryset.filter(company__slug=self.kwargs['company_slug'])


class SegmentsViewSet(ModelViewSet):
    serializer_class = SegmentSerializer
    queryset = Segment.objects.select_related('color__type', 'rack').all()
