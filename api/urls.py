from django.urls import include, path
from rest_framework import routers

from .views import SegmentsViewSet

router = routers.DefaultRouter()
router.register(r'segments', SegmentsViewSet)

urlpatterns = [
    path('', include(router.urls)),
    path('auth/', include('rest_framework.urls', namespace='rest_framework')),
]
