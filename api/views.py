from rest_framework.viewsets import ModelViewSet

from segments.models import Segment

from .serializers import SegmentSerializer


class SegmentsViewSet(ModelViewSet):
    serializer_class = SegmentSerializer
    queryset = Segment.objects.select_related('color__type', 'rack').all()
