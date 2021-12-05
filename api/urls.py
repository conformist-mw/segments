from django.urls import include, path
from rest_framework_nested import routers

from .views import SegmentsViewSet, CompanyViewSet, SectionViewSet

router = routers.SimpleRouter()
router.register('companies', CompanyViewSet)

companies_router = routers.NestedDefaultRouter(router, r'companies', lookup='company')
companies_router.register('sections', SectionViewSet)

urlpatterns = [
    path('', include(router.urls)),
    path('', include(companies_router.urls)),
    path('auth/', include('rest_framework.urls', namespace='rest_framework')),
]
